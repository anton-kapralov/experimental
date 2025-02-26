package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"golang.org/x/sync/errgroup"
)

const pqSerializationFailure = "40001"

func withTx(ctx context.Context, db *sqlx.DB, opts *sql.TxOptions, fn func(*sqlx.Tx) error) error {
	tx, err := db.BeginTxx(ctx, opts)
	if err != nil {
		return fmt.Errorf("failed to start a DB transaction: %w", err)
	}
	defer tx.Rollback()
	if err := fn(tx); err != nil {
		return err
	}
	return tx.Commit()
}

func withTxRetries(ctx context.Context, db *sqlx.DB, opts *sql.TxOptions, fn func(*sqlx.Tx) error) (int, error) {
	retries := 0
	for {
		err := withTx(ctx, db, opts, fn)
		if err == nil {
			return retries, nil
		}
		var pqErr *pq.Error
		if !errors.As(err, &pqErr) || pqErr.Code != pqSerializationFailure {
			return retries, err
		}
		retries++
	}
}

var schema = `
create table if not exists accounts
(
    id             bigint not null,
    balance_micros bigint not null default (0),
    constraint accounts_pk
        primary key (id)
);

create table if not exists transactions
(
    id uuid not null,
    at timestamptz not null,
    from_account_id bigint not null,
    to_account_id bigint not null,
    amount_micros bigint not null,
    constraint transactions_pk
        primary key (id),
    constraint transactions_from_account_id_fk
        foreign key (from_account_id) references accounts (id),
    constraint transactions_to_account_id_fk
        foreign key (to_account_id) references accounts (id)
);
`

const (
	cashAccountId = 1

	toMicros = 1_000_000
)

type repository struct {
	selectAccBalance  *sqlx.Stmt
	incrAccBalance    *sqlx.Stmt
	insertTransaction *sqlx.Stmt
}

func (r *repository) prepareStatements(db *sqlx.DB) error {
	for _, q := range []struct {
		stmt **sqlx.Stmt
		sql  string
	}{
		{
			&r.selectAccBalance,
			"select balance_micros from accounts where id=$1",
		},
		{
			&r.incrAccBalance,
			"update accounts set balance_micros=balance_micros+$1 where id=$2",
		},
		{
			&r.insertTransaction,
			"insert into transactions (id, at, from_account_id, to_account_id, amount_micros) values ($1, $2, $3, $4, $5)",
		},
	} {
		stmt, err := db.Preparex(q.sql)
		if err != nil {
			return fmt.Errorf("failed to prepare \"%s\": %v", q.sql, err)
		}
		*q.stmt = stmt
	}
	return nil
}

func (r *repository) getAccountBalance(tx *sqlx.Tx, id int64) (int64, error) {
	var balanceMicros int64
	if err := tx.Stmtx(r.selectAccBalance).Get(&balanceMicros, id); err != nil {
		return 0, err
	}
	return balanceMicros, nil
}

func (r *repository) increaseAccountBalance(tx *sqlx.Tx, id int64, amount int64) error {
	_, err := tx.Stmtx(r.incrAccBalance).Exec(amount, id)
	return err
}

func (r *repository) createTransaction(tx *sqlx.Tx, from, to int64, amount int64) error {
	_, err := tx.Stmtx(r.insertTransaction).Exec(uuid.NewString(), time.Now(), from, to, amount)
	return err
}

type service struct {
	repo *repository
}

func (s *service) deposit(tx *sqlx.Tx, to int64, amount int64) error {
	if err := s.repo.increaseAccountBalance(tx, cashAccountId, amount); err != nil {
		return fmt.Errorf("failed to increase acount balance for account id %d: %v", cashAccountId, err)
	}
	if err := s.repo.increaseAccountBalance(tx, to, amount); err != nil {
		return fmt.Errorf("failed to increase acount balance for account id %d: %v", to, err)
	}
	if err := s.repo.createTransaction(tx, cashAccountId, to, amount); err != nil {
		return fmt.Errorf("failed to create a transaction between account ids %d and %d: %v", cashAccountId, to, err)
	}
	return nil
}

func (s *service) pay(tx *sqlx.Tx, from, to int64, amount int64) error {
	log.Printf("pay(%d, %d, $%.2f): tx started", from, to, float64(amount)/toMicros)
	fromBalance, err := s.repo.getAccountBalance(tx, from)
	if err != nil {
		return fmt.Errorf("failed to get acount balance for account id %d: %v", from, err)
	}
	if fromBalance < amount {
		return errors.New("no sufficient funds")
	}
	log.Printf("pay(%d, %d, $%.2f): from balance $%.2f", from, to, float64(amount)/toMicros, float64(fromBalance)/toMicros)
	if err := s.repo.increaseAccountBalance(tx, from, -amount); err != nil {
		return fmt.Errorf("failed to decrease acount balance for account id %d: %w", from, err)
	}
	if err := s.repo.increaseAccountBalance(tx, to, amount); err != nil {
		return fmt.Errorf("failed to increase acount balance for account id %d: %w", to, err)
	}
	if err := s.repo.createTransaction(tx, from, to, amount); err != nil {
		return fmt.Errorf("failed to create a transaction between account ids %d and %d: %w", from, to, err)
	}

	log.Printf("pay(%d, %d, $%.2f): committing tx", from, to, float64(amount)/toMicros)
	return nil

}

func (s *service) balanceOf(tx *sqlx.Tx, accountID int64) (int64, error) {
	balance, err := s.repo.getAccountBalance(tx, accountID)
	if err != nil {
		return 0, fmt.Errorf("failed to get acount balance for account id %d: %w", accountID, err)
	}
	return balance, nil
}

type application struct {
	db      *sqlx.DB
	service *service
}

func (a *application) deposit(to int64, amount int64) error {
	txRetries, err := withTxRetries(context.TODO(), a.db, &sql.TxOptions{Isolation: sql.LevelSerializable}, func(tx *sqlx.Tx) error {
		return a.service.deposit(tx, to, amount)
	})
	if txRetries > 0 {
		log.Printf("deposit(%d, $%.2f): tx retried %d time(s)", to, float64(amount)/toMicros, txRetries)
	}
	return err
}

func (a *application) pay(from, to int64, amount int64) error {
	txRetries, err := withTxRetries(context.TODO(), a.db, &sql.TxOptions{Isolation: sql.LevelSerializable}, func(tx *sqlx.Tx) error {
		return a.service.pay(tx, from, to, amount)
	})
	if txRetries > 0 {
		log.Printf("pay(%d, %d, $%.2f): tx retried %d time(s)", from, to, float64(amount)/toMicros, txRetries)
	}
	return err
}

func (a *application) balanceOf(accountID int64) (int64, error) {
	var balance int64
	opts := &sql.TxOptions{Isolation: sql.LevelSerializable, ReadOnly: true}
	txRetries, err := withTxRetries(context.TODO(), a.db, opts, func(tx *sqlx.Tx) error {
		var err error
		balance, err = a.service.balanceOf(tx, accountID)
		return err
	})
	if txRetries > 0 {
		log.Printf("balanceOf(%d): tx retried %d time(s)", accountID, txRetries)
	}
	return balance, err
}

func transferUntilError(app *application, from, to int64, amountMicros int64) error {
	for {
		jitter := time.Duration(float64(time.Microsecond) * rand.Float64())
		time.Sleep(jitter)
		if err := app.pay(from, to, amountMicros); err != nil {
			return fmt.Errorf("pay(%d, %d, $%.2f): %w", from, to, float64(amountMicros)/toMicros, err)
		}
		fromBalance, err := app.balanceOf(from)
		if err != nil {
			return fmt.Errorf("balanceOf(%d): %w", from, err)
		}
		log.Printf(
			"pay(%d, %d, $%.2f): done; balance=$%.2f",
			from, to, float64(amountMicros)/toMicros, float64(fromBalance)/toMicros,
		)
	}
}

func main() {
	db, err := sqlx.Connect("postgres", "dbname=experimental sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)

	tx := db.MustBegin()
	tx.MustExec("insert into accounts (id) values (1), (1001), (1002), (1003) on conflict do nothing")
	if err := tx.Commit(); err != nil {
		log.Fatalln("failed to populate initial data")
	}

	repo := &repository{}
	if err := repo.prepareStatements(db); err != nil {
		log.Fatalln("failed to prepare statements", err)
	}
	app := &application{
		db: db,
		service: &service{
			repo: repo,
		},
	}

	const (
		aliceAccId   = 1001
		bobAccId     = 1002
		charlieAccId = 1003
	)
	if err := app.deposit(charlieAccId, 100*toMicros); err != nil {
		log.Fatalln(err)
	}
	var eg errgroup.Group
	eg.Go(func() error {
		return transferUntilError(app, charlieAccId, bobAccId, 9*toMicros)
	})
	eg.Go(func() error {
		return transferUntilError(app, charlieAccId, aliceAccId, 8*toMicros)
	})
	if err := eg.Wait(); err != nil {
		log.Println("Finished:", err)
	}
	charlieBalance, err := app.balanceOf(charlieAccId)
	if err != nil {
		log.Printf("balanceOf(%d): %s", charlieAccId, err)
	}
	log.Printf("Charlie's balance: $%.2f", float64(charlieBalance)/toMicros)
}

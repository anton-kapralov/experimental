package tetris

import (
	"sync"
	"time"
)

const (
	height = 20
	width  = 10
)

type Board [height][width]int

type GameOver struct{}

func (g GameOver) Error() string {
	return "Game Over"
}

type tetromino [4][2]int

func (t *tetromino) move(dr, dc int) {
	for i := range t {
		t[i][0] += dr
		t[i][1] += dc
	}
}

func (t *tetromino) rotate() {
	//TODO: implement me.
}

type Game struct {
	board   Board
	current tetromino
	mu      sync.Mutex
	Updated chan struct{}
}

func NewGame() *Game {
	g := &Game{
		Updated: make(chan struct{}),
	}
	_ = g.nextTetromino()
	return g
}

func (g *Game) Snapshot() Board {
	g.mu.Lock()
	defer g.mu.Unlock()
	var snapshot Board
	for i, row := range g.board {
		for j, v := range row {
			snapshot[i][j] = v
		}
	}
	for _, tile := range g.current {
		snapshot[tile[0]][tile[1]] = 1
	}
	return snapshot
}

type Command int

const (
	CommandUndefined Command = iota
	CommandUp
	CommandDown
	CommandRight
	CommandLeft
	CommandDrop
)

func (g *Game) Start(gameOver chan<- struct{}) {
	tick := time.Tick(1 * time.Second)
	for {
		<-tick
		if err := g.Move(CommandDown); err != nil {
			gameOver <- struct{}{}
			return
		}
	}
}

func (g *Game) Move(c Command) error {
	g.mu.Lock()
	defer g.mu.Unlock()
	defer g.onMoved()
	switch c {
	case CommandUp:
		return g.rotate()
	case CommandDown:
		return g.moveDown()
	case CommandRight:
		return g.moveHorizontally(1)
	case CommandLeft:
		return g.moveHorizontally(-1)
	case CommandDrop:
		return g.drop()
	}
	return nil
}

func (g *Game) onMoved() {
	g.Updated <- struct{}{}
}

func (g *Game) rotate() error {
	g.current.rotate()
	return nil
}

func (g *Game) moveDown() error {
	g.current.move(1, 0)
	if g.hasCollisions() {
		g.current.move(-1, 0)
		g.stickCurrent()
		return g.nextTetromino()
	}
	return nil
}

func (g *Game) moveHorizontally(dc int) error {
	g.current.move(0, dc)
	if g.hasCollisions() {
		g.current.move(0, -dc)
	}
	return nil
}

func (g *Game) drop() error {
	for !g.hasCollisions() {
		g.current.move(1, 0)
	}
	g.current.move(-1, 0)
	g.stickCurrent()
	return g.nextTetromino()
}

func (g *Game) stickCurrent() {
	for _, tile := range g.current {
		r := tile[0]
		c := tile[1]
		g.board[r][c] = 1
	}
}

func (g *Game) hasCollisions() bool {
	for _, tile := range g.current {
		r := tile[0]
		c := tile[1]
		if r < 0 || r >= height || c < 0 || c >= width || g.board[r][c] != 0 {
			return true
		}
	}
	return false
}

func (g *Game) nextTetromino() error {
	g.newTShapeTetromino()
	if g.hasCollisions() {
		return &GameOver{}
	}
	return nil
}

func (g *Game) newTShapeTetromino() {
	g.current[0][0] = 0
	g.current[0][1] = 1
	g.current[1][0] = 1
	g.current[1][1] = 0
	g.current[2][0] = 1
	g.current[2][1] = 1
	g.current[3][0] = 1
	g.current[3][1] = 2

	g.current.move(0, 4)
}

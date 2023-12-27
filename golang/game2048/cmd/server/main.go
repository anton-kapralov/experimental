package main

import (
	"flag"
	"fmt"
	"github.com/anton-kapralov/experimental/golang/game2048/internal/repository"
	"github.com/anton-kapralov/experimental/golang/game2048/internal/rest"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
)

type options struct {
	dbHost     string
	dbPort     int
	dbUser     string
	dbPassword string
	dbName     string
}

func (o *options) dbConnUri() string {
	dbCreds := ""
	if o.dbUser != "" {
		if o.dbPassword != "" {
			dbCreds = fmt.Sprintf("%s:%s@", o.dbUser, o.dbPassword)
		} else {
			dbCreds = fmt.Sprintf("%s@", o.dbUser)
		}
	}
	return fmt.Sprintf("mongodb://%s%s:%d", dbCreds, o.dbHost, o.dbPort)
}

func main() {
	opts := &options{}
	flag.StringVar(&opts.dbHost, "db-host", "localhost", "DB host")
	flag.IntVar(&opts.dbPort, "db-port", 27017, "DB port")
	flag.StringVar(&opts.dbUser, "db-user", "", "DB user")
	flag.StringVar(&opts.dbPassword, "db-password", "", "DB password")
	flag.StringVar(&opts.dbName, "db-name", "game2048", "DB name (schema)")
	flag.Parse()

	repo, err := repository.New(context.TODO(), opts.dbConnUri(), opts.dbName)
	if err != nil {
		log.Fatalln(err)
	}

	restServer := rest.NewServer(repo)

	router := gin.Default()
	router.GET("/", restServer.Index)
	router.POST("/games", restServer.NewGame)
	router.GET("/games/:key", restServer.GetGame)
	router.POST("/games/:key", restServer.MoveGame)

	addr := ":8080"
	log.Println("Now listening on", addr)
	log.Fatal(router.Run(addr))
}

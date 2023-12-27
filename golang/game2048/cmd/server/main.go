package main

import (
	"github.com/anton-kapralov/experimental/golang/game2048/internal/rest"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	restServer := &rest.Server{}

	router := gin.Default()
	router.GET("/", restServer.Index)
	router.POST("/games", restServer.NewGame)
	router.GET("/games/:key", restServer.GetGame)
	router.POST("/games/:key", restServer.MoveGame)

	addr := ":8080"
	log.Println("Now listening on", addr)
	log.Fatal(router.Run(addr))
}

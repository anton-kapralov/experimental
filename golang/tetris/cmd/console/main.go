package main

import (
	"fmt"
	"github.com/anton-kapralov/experimental/golang/console"
	"github.com/anton-kapralov/experimental/golang/tetris"
	"log"
)

var (
	gameOver = make(chan struct{})
)

func main() {
	c, err := console.New()
	if err != nil {
		panic(err)
	}
	defer c.Close()
	c.Clear()

	game := tetris.NewGame()
	draw(c, game)

	go inputLoop(c, game)
	go onUpdated(c, game)
	go game.Start(gameOver)

	<-gameOver
}

func draw(c *console.Console, game *tetris.Game) {
	c.MoveCursor(0, 0)
	fmt.Print(" ")
	board := game.Snapshot()
	for i := 0; i < len(board[0]); i++ {
		fmt.Print("_")
	}
	fmt.Print("\r\n")

	for i, row := range board {
		fmt.Print("|")
		for _, v := range row {
			if v == 1 {
				fmt.Print("≣")
				continue
			}
			if i == 19 {
				fmt.Print("_")
				continue
			}
			fmt.Print(".")
		}
		fmt.Print("|\r\n")
	}
}

func inputLoop(c *console.Console, game *tetris.Game) {
loop:
	for {
		code, err := c.Read()
		if err != nil {
			log.Fatal(err)
		}
		switch code {
		case console.Up:
			fmt.Printf("⬆️\r\n")
			if err := game.Move(tetris.CommandRotateClockwise); err != nil {
				break loop
			}
		case console.Alt | console.Up:
			fmt.Printf("⇧\r\n")
			if err := game.Move(tetris.CommandRotateCounterClockwise); err != nil {
				break loop
			}
		case console.Down:
			fmt.Printf("⬇️\r\n")
			if err := game.Move(tetris.CommandDrop); err != nil {
				break loop
			}
		case console.Right:
			fmt.Printf("➡️\r\n")
			if err := game.Move(tetris.CommandRight); err != nil {
				break loop
			}
		case console.Left:
			fmt.Printf("⬅️\r\n")
			if err := game.Move(tetris.CommandLeft); err != nil {
				break loop
			}
		case console.Sigint:
			break loop
		default:
			fmt.Printf("%x\r\n", code)
		}
	}

	gameOver <- struct{}{}
}

func onUpdated(c *console.Console, game *tetris.Game) {
	for {
		<-game.Updated
		draw(c, game)
	}
}

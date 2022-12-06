package main

import (
	"fmt"
	"github.com/anton-kapralov/experimental/golang/console"
	"log"
)

const (
	height = 20
	width  = 10
)

type tetrisBoard [height][width]int
type tetromino [4][2]int

func (t *tetromino) move(dr, dc int) {
	for i := range t {
		t[i][0] += dr
		t[i][1] += dc
	}
}

var (
	board   tetrisBoard
	current tetromino
)

func main() {
	c, err := console.New()
	if err != nil {
		panic(err)
	}
	defer c.Close()
	c.Clear()

	nextTetromino()

inputLoop:
	for {
		c.MoveCursor(0, 0)
		draw()
		code, err := c.Read()
		if err != nil {
			log.Fatal(err)
		}
		switch code {
		case console.Up:
			fmt.Printf("⬆️\r\n")
		case console.Down:
			fmt.Printf("⬇️\r\n")
		case console.Right:
			fmt.Printf("➡️\r\n")
		case console.Left:
			fmt.Printf("⬅️\r\n")
		case console.Sigint:
			break inputLoop
		default:
			fmt.Printf("%x\r\n", code)
		}
	}
}

func nextTetromino() {
	newTShapeTetromino()
}

func draw() {
	fmt.Print(" ")
	for i := 0; i < width; i++ {
		fmt.Print("_")
	}
	fmt.Print("\r\n")

	b := getSnapshot()
	for i, row := range b {
		fmt.Print("|")
		for _, v := range row {
			if v == 1 {
				fmt.Print("⚄")
				continue
			}
			if i == 19 {
				fmt.Print("_")
				continue
			}
			fmt.Print(" ")
		}
		fmt.Print("|\r\n")
	}
}

func getSnapshot() tetrisBoard {
	var snapshot tetrisBoard
	for i, row := range board {
		for j, v := range row {
			snapshot[i][j] = v
		}
	}
	for _, tile := range current {
		snapshot[tile[0]][tile[1]] = 1
	}
	return snapshot
}

func newTShapeTetromino() {
	current[0][0] = 0
	current[0][1] = 1
	current[1][0] = 1
	current[1][1] = 0
	current[2][0] = 1
	current[2][1] = 1
	current[3][0] = 1
	current[3][1] = 2

	current.move(0, 4)
}

package main

import (
	"fmt"
	"github.com/anton-kapralov/experimental/golang/console"
	"log"
	"sync"
	"time"
)

const (
	height = 20
	width  = 10
)

type command int

const (
	commandUndefined command = iota
	commandUp
	commandDown
	commandRight
	commandLeft
	commandDrop
)

type tetrisBoard [height][width]int
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

var (
	board    tetrisBoard
	current  tetromino
	mu       sync.Mutex
	gameOver = make(chan struct{})
)

func getSnapshot() tetrisBoard {
	mu.Lock()
	defer mu.Unlock()
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

func main() {
	c, err := console.New()
	if err != nil {
		panic(err)
	}
	defer c.Close()
	c.Clear()

	commands := make(chan command)

	nextTetromino()
	draw(c)

	go gameLoop(commands)
	go handleCommands(commands, func() { draw(c) })
	go inputLoop(c, commands)

	<-gameOver
}

func inputLoop(c *console.Console, commands chan command) {
	for {
		code, err := c.Read()
		if err != nil {
			log.Fatal(err)
		}
		switch code {
		case console.Up:
			fmt.Printf("⬆️\r\n")
			commands <- commandUp
		case console.Down:
			fmt.Printf("⬇️\r\n")
			commands <- commandDrop
		case console.Right:
			fmt.Printf("➡️\r\n")
			commands <- commandRight
		case console.Left:
			fmt.Printf("⬅️\r\n")
			commands <- commandLeft
		case console.Sigint:
			gameOver <- struct{}{}
			return
		default:
			fmt.Printf("%x\r\n", code)
		}
	}
}

func draw(c *console.Console) {
	c.MoveCursor(0, 0)
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

func nextTetromino() {
	newTShapeTetromino()
	if hasCollisions() {
		gameOver <- struct{}{}
	}
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

func gameLoop(commands chan<- command) {
	tick := time.Tick(1 * time.Second)
	for {
		<-tick
		commands <- commandDown
	}
}

func handleCommands(commands <-chan command, onStateChanged func()) {
	for {
		c := <-commands
		if c == commandUndefined {
			return
		}
		move(c)
		onStateChanged()
	}
}

func move(c command) {
	mu.Lock()
	defer mu.Unlock()
	switch c {
	case commandUp:
		rotate()
	case commandDown:
		moveDown()
	case commandDrop:
		drop()
	}
}

func rotate() {
	current.rotate()
}

func moveDown() {
	current.move(1, 0)
	if hasCollisions() {
		current.move(-1, 0)
		stickCurrent()
		nextTetromino()
	}
}

func drop() {
	for !hasCollisions() {
		current.move(1, 0)
	}
	current.move(-1, 0)
	stickCurrent()
	nextTetromino()
}

func hasCollisions() bool {
	for _, tile := range current {
		r := tile[0]
		c := tile[1]
		if r < 0 || r >= height || c < 0 || c >= width || board[r][c] != 0 {
			return true
		}
	}
	return false
}

func stickCurrent() {
	for _, tile := range current {
		r := tile[0]
		c := tile[1]
		board[r][c] = 1
	}
}

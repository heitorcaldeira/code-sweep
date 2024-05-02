package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	mines "github.com/heitorcaldeira/code-sweep/pkg"
)

func TerminalPlay() {
	board := mines.NewMineBoard(10, 10, 10)
	timer := time.NewTicker(100 * time.Millisecond)

	for {
		<-timer.C
		fmt.Print("\033[H\033[2J")
		board.Debug(false)
		fmt.Println("Insert a row and a col: (XY)")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		cmd := strings.Split(line, "")
		row, err := strconv.ParseFloat(cmd[0], 64)

		if err != nil {
			log.Fatal(err)
		}

		col, err := strconv.ParseFloat(cmd[1], 64)

		if err != nil {
			log.Fatal(err)
		}

		err = board.PickCell(int(row), int(col))

		if err != nil {
			log.Fatal(err)
		}

		board.Debug(false)

		if board.Status == mines.GameOver {
			log.Fatal("GAME OVER")
		}

		board.CheckForWin()

		if board.Status == mines.Victory {
			log.Fatal("YOU WON")
		}
	}
}

func main() {
	TerminalPlay()
}

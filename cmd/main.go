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
  board := mines.NewMineBoard(10, 5, 0)
  board.Debug(false)

  timer := time.NewTicker(100 * time.Millisecond)

  for {
    <- timer.C
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
  }
}

func main() {
  TerminalPlay()
}

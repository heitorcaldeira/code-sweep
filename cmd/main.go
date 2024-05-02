package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	mines "github.com/heitorcaldeira/code-sweep/pkg"
)

func main() {
  board := mines.NewMineBoard(10, 10, 5)
  board.Debug(true)

  timer := time.NewTicker(100 * time.Millisecond)

  for {
    <- timer.C
    // err := board.PickCell(int(math.Floor(rand.Float64() * 10)), int(math.Floor(rand.Float64() * 10)))
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

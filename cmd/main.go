package main

import (
	"log"
	"math"
	"math/rand"
	"time"

	mines "github.com/heitorcaldeira/code-sweep/pkg"
)

func main() {
  board := mines.NewMineBoard(10, 10, 10)
  board.Debug(true)

  timer := time.NewTicker(2000 * time.Millisecond)

  for {
    <- timer.C
    err := board.PickCell(int(math.Floor(rand.Float64() * 10)), int(math.Floor(rand.Float64() * 10)))

    if err != nil {
      log.Fatal(err)
    }

    board.Debug(false)

    if board.Status == mines.GameOver {
      log.Fatal("GAME OVER")
    }
  }
}

package main

import mines "github.com/heitorcaldeira/code-sweep/pkg"

func main() {
  board := mines.NewMineBoard(10, 10, 5)
  board.Debug()
}

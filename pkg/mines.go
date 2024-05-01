package mines

import (
	"fmt"
	"math"
	"math/rand"
)

type CellState = int

const (
  Closed CellState = 0
  Opened CellState = 1
)

type MineCell struct {
  Row int
  Col int
  State CellState
  Content int
}

type MineBoard struct {
  Cells []MineCell
  Rows int
  Cols int
  Bombs int
}

func (b *MineBoard) CreateEmptyBoard() {
  cells := make([]MineCell, 0)

  for row := 0; row < b.Rows; row++ {
    for col := 0; col < b.Cols; col++ {
      cells = append(cells, MineCell{Row: row, Col: col, State: Closed})
    }
  }

  b.Cells = cells
}

func (b *MineBoard) AddBombRandomCell() {
  item := int32(math.Floor(rand.Float64() * float64(len(b.Cells))))

  if b.Cells[item].Content == 0 {
    b.Cells[item].Content = -1
  } else {
    b.AddBombRandomCell()
  }
}

func (b *MineBoard) CreateBombs() {
  for i := 0; i < b.Bombs; i++ {
    b.AddBombRandomCell()
  }
}

func (b *MineBoard) Debug() {
  output := ""
  count := 0
  for row := 0; row < b.Rows; row++ {
    for col := 0; col < b.Cols; col++ {
      var content = b.Cells[count].Content
      if content == 0 {
        output += " "
      } else if content > 0 {
        output += fmt.Sprint(content)
      } else {
        output += "X"
      }

      count += 1
    }
    output += "\n"
  }

  fmt.Println(output)
}

func NewMineBoard(rows, cols, totalBombs int) MineBoard {
  board := MineBoard{Rows: rows, Cols: cols, Bombs: totalBombs}
  board.CreateEmptyBoard()
  board.CreateBombs()

  return board
}

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

type GameStatus = int

const (
	Iddle    GameStatus = 0
	Running  GameStatus = 1
	GameOver GameStatus = 2
)

type MineCell struct {
	Row     int
	Col     int
	State   CellState
	Content int
}

type MineBoard struct {
	Cells  [][]MineCell
	Rows   int
	Cols   int
	Bombs  int
	Status GameStatus
}

func (b *MineBoard) CreateEmptyBoard() {
	cells := make([][]MineCell, b.Rows)
	for row := 0; row < b.Rows; row++ {
		for col := 0; col < b.Cols; col++ {
			cells[row] = append(cells[row], MineCell{Row: row, Col: col, State: Closed})
		}
	}

	b.Cells = cells
}

func (b *MineBoard) AddBombRandomCell() {
	row := int32(math.Floor(rand.Float64() * float64(b.Rows)))
	col := int32(math.Floor(rand.Float64() * float64(b.Cols)))

	if b.Cells[row][col].Content == 0 {
		b.Cells[row][col].Content = -1
	} else {
		b.AddBombRandomCell()
	}
}

func (b *MineBoard) CreateBombs() {
	for i := 0; i < b.Bombs; i++ {
		b.AddBombRandomCell()
	}
}

func (b *MineBoard) Debug(showAll bool) {
	output := ""
	count := 0
	for _, row := range b.Cells {
		for _, col := range row {
			var content = col.Content

			if !showAll && col.State == Closed {
				output += "▒"
				continue
			}

			if content == 0 {
				output += "o"
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

func (b *MineBoard) FillSmartCells() {
	for idxRow, row := range b.Cells {
		for idxCol, col := range row {
			currentRow := 0
			currentCol := 0
			content := 0

			if col.Content == -1 {
				continue
			}

			// corner top left
			currentRow = idxRow - 1
			currentCol = idxCol - 1
			if currentRow >= 0 && currentCol >= 0 && b.Cells[currentRow][currentCol].Content == -1 {
				content += 1
			}

			// top center
			currentRow = idxRow - 1
			currentCol = idxCol
			if currentRow >= 0 && b.Cells[currentRow][currentCol].Content == -1 {
				content += 1
			}

			// corner top right
			currentRow = idxRow - 1
			currentCol = idxCol + 1
			if currentRow >= 0 && currentCol < b.Cols && b.Cells[currentRow][currentCol].Content == -1 {
				content += 1
			}

			// right center
			currentRow = idxRow
			currentCol = idxCol + 1
			if currentCol < b.Cols && b.Cells[currentRow][currentCol].Content == -1 {
				content += 1
			}

			// corner bottom right
			currentRow = idxRow + 1
			currentCol = idxCol + 1
			if currentRow < b.Rows && currentCol < b.Cols && b.Cells[currentRow][currentCol].Content == -1 {
				content += 1
			}

			// bottom center
			currentRow = idxRow + 1
			currentCol = idxCol
			if currentRow < b.Rows && b.Cells[currentRow][currentCol].Content == -1 {
				content += 1
			}

			// corner bottom left
			currentRow = idxRow + 1
			currentCol = idxCol - 1
			if currentRow < b.Rows && currentCol >= 0 && b.Cells[currentRow][currentCol].Content == -1 {
				content += 1
			}

			// center right
			currentRow = idxRow
			currentCol = idxCol - 1
			if currentCol >= 0 && b.Cells[currentRow][currentCol].Content == -1 {
				content += 1
			}

			b.Cells[idxRow][idxCol].Content = content
		}
	}
}

func (b *MineBoard) OpenBlankCells(row, col int) {
	var idxs = [][]int{
    {-1, -1}, // top left
    {-1, 0}, // top center
    {-1, 1}, // top right
    {0, 1}, // right center
    {1, 1}, // bottom right
    {1, 0}, // bottom center
    {1, -1}, //bottom left
    {0, -1}, // left center
  }

  for _, idx := range idxs {
    currentRow := row + idx[0]
    currentCol := col + idx[1]
    if currentRow >= 0 && currentRow < b.Rows && currentCol >= 0 && currentCol < b.Cols {
      b.Cells[row][col].State = Opened

      content := b.Cells[currentRow][currentCol].Content 

      if content > 0 {
        b.Cells[currentRow][currentCol].State = Opened
      } else if content == 0 && b.Cells[currentRow][currentCol].State == Closed {
        b.OpenBlankCells(currentRow, currentCol)
      }
    }
	}
}

func (p *MineBoard) PickCell(row, col int) error {
  fmt.Printf("row %d col %d", row, col)
  fmt.Println()
	if row >= 0 && row < p.Rows && col >= 0 && col < p.Cols {
		var content = p.Cells[row][col].Content

		switch content {
		case -1:
			p.Status = GameOver
		case 0:
			p.OpenBlankCells(row, col)
		}

		p.Cells[row][col].State = Opened

		return nil
	}

	return fmt.Errorf("invalid placement")
}

func NewMineBoard(rows, cols, totalBombs int) MineBoard {
	board := MineBoard{Rows: rows, Cols: cols, Bombs: totalBombs, Status: Iddle}
	board.CreateEmptyBoard()
	board.CreateBombs()
	board.FillSmartCells()

	return board
}

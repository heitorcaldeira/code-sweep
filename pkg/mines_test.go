package mines

import "testing"

func TestCreateEmptyBoard(t *testing.T) {
  var rows = 10
  var cols = 5
  var totalBombs = 5
  board := MineBoard{Rows: rows, Cols: cols, Bombs: totalBombs}
  board.CreateEmptyBoard()

  var cells = board.Cells

  if len(cells) != rows || len(cells[0]) != cols {
    t.Errorf("rows should have length %d, but got %d", rows, len(cells))
  }

  if len(cells[0]) != cols {
    t.Errorf("cols should have length %d, but got %d", cols, len(cells[0]))
  }

  var row = cells[rows - 1]
  var col = row[cols - 1]

  if col.Row != rows - 1 || col.Col != cols -1 || col.State != 0 || col.Content != 0 {
    t.Errorf("last cell does not have correct values for row and col %+v", cells)
  }
}

func TestCreateBombs(t *testing.T) {
  var rows = 10
  var cols = 5
  var totalBombs = 5
  board := MineBoard{Rows: rows, Cols: cols, Bombs: totalBombs}
  board.CreateEmptyBoard()
  board.CreateBombs()
  countBombs := 0

  for _, row := range board.Cells {
    for _, col := range row {
      if col.Content == -1 {
        countBombs += 1
      }
    }
  }

  if countBombs != board.Bombs {
    t.Errorf("number of bombs should be %d, received %d", board.Bombs, countBombs)
  }
}


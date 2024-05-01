package mines

import "testing"

func TestNewMineBoard(t *testing.T) {
  board := NewMineBoard(10, 5, 5)
  var cells = board.Cells

  if len(cells) != 50 {
    t.Errorf("cells should have length 50, but got %d", len(cells))
  }

  var cell = cells[49]

  if cell.Row != 9 || cell.Col != 4 || cell.State != 0 || cell.Content != 0 {
    t.Errorf("last cell does not have correct values for row and col %+v", cells)
  }

  countBombs := 0

  for _, cell := range cells {
    if cell.Content == -1 {
      countBombs += 1
    }
  }

  if countBombs != board.Bombs {
    t.Errorf("number of bombs should be %d, received %d", board.Bombs, countBombs)
  }
}


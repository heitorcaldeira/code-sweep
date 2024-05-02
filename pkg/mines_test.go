package mines

import "testing"

var rows = 10
var cols = 5
var totalBombs = 5

func TestCreateEmptyBoard(t *testing.T) {
	board := MineBoard{Rows: rows, Cols: cols, Bombs: totalBombs}
	board.CreateEmptyBoard()

	var cells = board.Cells

	if len(cells) != rows || len(cells[0]) != cols {
		t.Errorf("rows should have length %d, but got %d", rows, len(cells))
	}

	if len(cells[0]) != cols {
		t.Errorf("cols should have length %d, but got %d", cols, len(cells[0]))
	}

	var row = cells[rows-1]
	var col = row[cols-1]

	if col.Row != rows-1 || col.Col != cols-1 || col.State != 0 || col.Content != 0 {
		t.Errorf("last cell does not have correct values for row and col %+v", cells)
	}
}

func TestCreateBombs(t *testing.T) {
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

func TestPickCell(t *testing.T) {
	board := MineBoard{Rows: rows, Cols: cols, Bombs: 50, Status: Running}
	board.CreateEmptyBoard()
	board.CreateBombs()
	board.FillSmartCells()
	board.PickCell(4, 4)

	if board.Status != GameOver {
		t.Errorf("picked a bomb and game is not over")
	}

	board = MineBoard{Rows: rows, Cols: cols, Bombs: 0, Status: Running}
	board.CreateEmptyBoard()
	board.FillSmartCells()
	board.PickCell(4, 4)

	if board.Cells[4][4].State != Opened {
		t.Errorf("the cell should be opened, but is closed")
	}
}

func TestOpenBlankCells(t *testing.T) {
	board := MineBoard{Rows: rows, Cols: cols, Bombs: 0, Status: Running}
	board.CreateEmptyBoard()
	board.CreateBombs()
	board.FillSmartCells()
	board.PickCell(4, 4)
	board.CheckForWin()

	if board.Status != Victory {
		t.Errorf("the game should return a victory status, but returned %d", board.Status)
	}
}

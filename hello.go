package main

import (
	"github.com/rivo/tview"
	// "fmt"
)

type Piece uint8

const (
	NAME string = "Lichess TUI"
	CELL_SIZE int = 8
)

const (
	Pawn   Piece = 0
	Rook   Piece = 1
	Knight Piece = 2
	Bishop Piece = 3
	Queen  Piece = 4
	King   Piece = 5
)

func newPrimitive(text string) tview.Primitive {
	return tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText(text)
}
func main() {
	menu := newPrimitive("Menu")
	// main := newPrimitive("Main content")
	// sideBar := newPrimitive("Side Bar")
	//
	grid := tview.NewGrid().
		SetSize(8, 8, 8 * CELL_SIZE, 8 * CELL_SIZE).
		SetBorders(true)
		// AddItem(newPrimitive("Header"), 0, 0, 1, 3, 0, 0, false).
		// AddItem(newPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)

	for i := range 8 {
		for j := range 8 {
			grid.AddItem(menu, i, j, i * CELL_SIZE, j * CELL_SIZE, 0, 0, true)
		}
	}

	if err := tview.NewApplication().SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

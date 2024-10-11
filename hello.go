package main

import (
	"github.com/rivo/tview"
)

type Piece uint8

const (
	Pawn Piece = 0
	Rook Piece = 1
	Knight Piece = 2
	Bishop Piece = 3
	Queen Piece = 4
	King Piece = 5
)

func main() {
	board := make([][]Piece, 8 * 8)

	box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}

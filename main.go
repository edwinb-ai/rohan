package main

import (
	"fmt"

	"github.com/edwinb-ai/chess-engine/board"
)

func main() {
	b := board.NewChessBoard()

	fmt.Println("Initial Board Setup, lowercase is Black:")
	b.PrintBoard()
}

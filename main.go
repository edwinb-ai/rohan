package main

import (
	"fmt"

	"github.com/edwinb-ai/chess-engine/board"
)

func main() {
	cb := board.NewChessBoard()
	fen := cb.ToFEN("w", "KQkq", "-", 0, 1)
	fmt.Println(fen)

	fmt.Println("Initial Board Setup, lowercase is Black:")
	cb.PrintBoard()
}

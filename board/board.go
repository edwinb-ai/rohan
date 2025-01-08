package board

import (
	"fmt"
)

type Bitboard uint64

type ChessBoard struct {
	WhitePawns   Bitboard
	WhiteKnights Bitboard
	WhiteBishops Bitboard
	WhiteRooks   Bitboard
	WhiteQueens  Bitboard
	WhiteKing    Bitboard
	BlackPawns   Bitboard
	BlackKnights Bitboard
	BlackBishops Bitboard
	BlackRooks   Bitboard
	BlackQueens  Bitboard
	BlackKing    Bitboard
}

// Initialize the chessboard with the starting position
func NewChessBoard() ChessBoard {
	return ChessBoard{
		WhitePawns:   0x000000000000FF00, // Rank 2
		WhiteRooks:   0x0000000000000081, // a1 and h1
		WhiteKnights: 0x0000000000000042, // b1 and g1
		WhiteBishops: 0x0000000000000024, // c1 and f1
		WhiteQueens:  0x0000000000000008, // d1
		WhiteKing:    0x0000000000000010, // e1
		BlackPawns:   0x00FF000000000000, // Rank 7
		BlackRooks:   0x8100000000000000, // a8 and h8
		BlackKnights: 0x4200000000000000, // b8 and g8
		BlackBishops: 0x2400000000000000, // c8 and f8
		BlackQueens:  0x0800000000000000, // d8
		BlackKing:    0x1000000000000000, // e8
	}
}

// Print the bitboard representation for debugging
func printBitboard(bb Bitboard) {
	for rank := 7; rank >= 0; rank-- {
		for file := 0; file < 8; file++ {
			square := rank*8 + file
			if (bb & (1 << square)) != 0 {
				fmt.Print("1 ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (cb ChessBoard) PrintBoard() {
	// Define symbols for pieces
	symbols := map[string]string{
		"WhitePawns":   "P",
		"WhiteKnights": "N",
		"WhiteBishops": "B",
		"WhiteRooks":   "R",
		"WhiteQueens":  "Q",
		"WhiteKing":    "K",
		"BlackPawns":   "p",
		"BlackKnights": "n",
		"BlackBishops": "b",
		"BlackRooks":   "r",
		"BlackQueens":  "q",
		"BlackKing":    "k",
	}

	// Loop through each square on the board
	for rank := 7; rank >= 0; rank-- {
		for file := 0; file < 8; file++ {
			square := rank*8 + file
			found := false

			// Check each piece's bitboard to determine the piece on this square
			for pieceName, bitboard := range map[string]Bitboard{
				"WhitePawns":   cb.WhitePawns,
				"WhiteKnights": cb.WhiteKnights,
				"WhiteBishops": cb.WhiteBishops,
				"WhiteRooks":   cb.WhiteRooks,
				"WhiteQueens":  cb.WhiteQueens,
				"WhiteKing":    cb.WhiteKing,
				"BlackPawns":   cb.BlackPawns,
				"BlackKnights": cb.BlackKnights,
				"BlackBishops": cb.BlackBishops,
				"BlackRooks":   cb.BlackRooks,
				"BlackQueens":  cb.BlackQueens,
				"BlackKing":    cb.BlackKing,
			} {
				if isBitSet(bitboard, square) {
					fmt.Print(symbols[pieceName], " ")
					found = true
					break
				}
			}

			// If no piece is found, print a dot for an empty square
			if !found {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

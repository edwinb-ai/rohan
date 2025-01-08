package board

import (
	"fmt"
	"strconv"
	"strings"
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

// A function to take the state of the board and create a FEN string
func (cb ChessBoard) ToFEN(activePlayer string, castlingRights string, enPassant string, halfMoveClock int, fullMoveNumber int) string {
	var fen strings.Builder

	// Combine all bitboards to detect occupied squares
	allPieces := map[string]Bitboard{
		"P": cb.WhitePawns,
		"N": cb.WhiteKnights,
		"B": cb.WhiteBishops,
		"R": cb.WhiteRooks,
		"Q": cb.WhiteQueens,
		"K": cb.WhiteKing,
		"p": cb.BlackPawns,
		"n": cb.BlackKnights,
		"b": cb.BlackBishops,
		"r": cb.BlackRooks,
		"q": cb.BlackQueens,
		"k": cb.BlackKing,
	}

	// Piece Placement
	for rank := 7; rank >= 0; rank-- {
		emptyCount := 0
		for file := 0; file < 8; file++ {
			square := rank*8 + file
			found := false

			// Check if the square is occupied by a piece
			for pieceName, bitboard := range allPieces {
				if (bitboard & (1 << square)) != 0 {
					// If empty squares preceded, add the count to the FEN
					if emptyCount > 0 {
						fen.WriteString(strconv.Itoa(emptyCount))
						emptyCount = 0
					}
					fen.WriteString(pieceName)
					found = true
					break
				}
			}

			// Count empty squares
			if !found {
				emptyCount++
			}
		}

		// Add any trailing empty squares
		if emptyCount > 0 {
			fen.WriteString(strconv.Itoa(emptyCount))
		}

		// Add a slash after every rank except the last
		if rank > 0 {
			fen.WriteString("/")
		}
	}

	// Active player
	fen.WriteString(" ")
	fen.WriteString(activePlayer)

	// Castling rights
	fen.WriteString(" ")
	if castlingRights == "" {
		fen.WriteString("-")
	} else {
		fen.WriteString(castlingRights)
	}

	// En passant target square
	fen.WriteString(" ")
	if enPassant == "" {
		fen.WriteString("-")
	} else {
		fen.WriteString(enPassant)
	}

	// Half-move clock
	fen.WriteString(" ")
	fen.WriteString(strconv.Itoa(halfMoveClock))

	// Full move number
	fen.WriteString(" ")
	fen.WriteString(strconv.Itoa(fullMoveNumber))

	return fen.String()
}

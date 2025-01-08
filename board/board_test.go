package board

import (
	"testing"
)

func TestNewChessBoard(t *testing.T) {
	board := NewChessBoard()

	// Check the initial position of white pawns
	expectedWhitePawns := Bitboard(0x000000000000FF00)
	if board.WhitePawns != expectedWhitePawns {
		t.Errorf("Expected white pawns bitboard: %064b, got: %064b", expectedWhitePawns, board.WhitePawns)
	}

	// Check the initial position of black rooks
	expectedBlackRooks := Bitboard(0x8100000000000000)
	if board.BlackRooks != expectedBlackRooks {
		t.Errorf("Expected black rooks bitboard: %064b, got: %064b", expectedBlackRooks, board.BlackRooks)
	}
}

func TestBitManipulation(t *testing.T) {
	var bb Bitboard

	// Test setting a bit
	bb = setBit(bb, 0)
	if !isBitSet(bb, 0) {
		t.Errorf("Bit 0 should be set, but it's not")
	}

	// Test clearing a bit
	bb = clearBit(bb, 0)
	if isBitSet(bb, 0) {
		t.Errorf("Bit 0 should be cleared, but it's not")
	}

	// Test squareToBitboard
	expected := Bitboard(1)
	if squareToBitboard(0) != expected {
		t.Errorf("Expected bitboard for square 0: %064b, got: %064b", expected, squareToBitboard(0))
	}
}

func TestToFEN(t *testing.T) {
	cb := NewChessBoard()

	fen := cb.ToFEN("w", "KQkq", "-", 0, 1)
	expectedFEN := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

	if fen != expectedFEN {
		t.Errorf("Expected FEN: %s, got: %s", expectedFEN, fen)
	}
}

func TestGeneratePawnMoves(t *testing.T) {
	cb := NewChessBoard()

	// Test White Pawn Moves
	occupied := cb.OccupiedSquares()
	opponent := cb.OpponentPieces(true) // White is moving
	moves := generatePawnMoves(cb.WhitePawns, occupied, opponent, true)

	// Expected: White pawns can move forward one square from rank 2
	expected := Bitboard(0x0000000000FF0000) // Rank 3
	if moves != expected {
		t.Errorf("Expected white pawn moves: %064b, got: %064b", expected, moves)
	}

	// Test Black Pawn Moves
	opponent = cb.OpponentPieces(false) // Black is moving
	moves = generatePawnMoves(cb.BlackPawns, occupied, opponent, false)

	// Expected: Black pawns can move forward one square from rank 7
	expected = Bitboard(0x0000FF0000000000) // Rank 6
	if moves != expected {
		t.Errorf("Expected black pawn moves: %064b, got: %064b", expected, moves)
	}
}

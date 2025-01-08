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

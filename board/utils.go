package board

// Set a bit at a specific square
func setBit(bb Bitboard, square int) Bitboard {
	return bb | (1 << square)
}

// Clear a bit at a specific square
func clearBit(bb Bitboard, square int) Bitboard {
	return bb & ^(1 << square)
}

// Check if a bit is set
func isBitSet(bb Bitboard, square int) bool {
	return (bb & (1 << square)) != 0
}

// Convert a square index (0-63) to a bitboard with only that bit set
func squareToBitboard(square int) Bitboard {
	return 1 << square
}

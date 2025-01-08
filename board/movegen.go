package board

func generatePawnMoves(pawns Bitboard, occupied Bitboard, opponent Bitboard, isWhite bool) Bitboard {
	var moves Bitboard
	emptySquares := ^occupied

	if isWhite {
		// Single forward moves: pawns move forward to empty squares
		moves |= (pawns << 8) & emptySquares

		// Double forward moves: only for pawns on rank 2, and both squares must be empty
		doubleMoves := (pawns & 0x000000000000FF00) << 16
		moves |= doubleMoves & emptySquares & (emptySquares << 8)

		// Captures: diagonal left and right
		moves |= (pawns << 7) & opponent & ^Bitboard(0x0101010101010101) // Exclude a-file
		moves |= (pawns << 9) & opponent & ^Bitboard(0x8080808080808080) // Exclude h-file
	} else {
		// Single forward moves: pawns move forward to empty squares
		moves |= (pawns >> 8) & emptySquares

		// Double forward moves: only for pawns on rank 7, and both squares must be empty
		doubleMoves := (pawns & 0x00FF000000000000) >> 16
		moves |= doubleMoves & emptySquares & (emptySquares >> 8)

		// Captures: diagonal left and right
		moves |= (pawns >> 7) & opponent & ^Bitboard(0x8080808080808080) // Exclude h-file
		moves |= (pawns >> 9) & opponent & ^Bitboard(0x0101010101010101) // Exclude a-file
	}

	return moves
}

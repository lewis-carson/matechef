package matechef

import (
	"math/bits"

	"github.com/notnil/chess"
)

//"fmt"

var centerSquares = []chess.Square{
	chess.E4,
	chess.D4,
	chess.E5,
	chess.D5,
}

var castleSquares = []chess.Square{
	chess.C1,
	chess.C8,
	chess.G1,
	chess.G8,
}

var weights = map[chess.PieceType]int{
	chess.Queen:  950,
	chess.Bishop: 325,
	chess.Knight: 305,
	chess.Pawn:   100,
	chess.King:   0,
	chess.Rook:   563,
}

func evalKingSafety(board *chess.Board, position *chess.Position) int {
	colour := position.Turn()
	kingSq := board.PublicBB().BlackKingSq
	if colour == chess.White {
		kingSq = board.PublicBB().WhiteKingSq
	}
	kqBB := uint64(chess.PossibleMoves(position, chess.Queen, kingSq))
	return -bits.OnesCount64(kqBB) * 20
}

func evalCenterSquares(board *chess.Board, position *chess.Position) int {
	eval := 0
	for _, square := range centerSquares {
		if board.Piece(square).Color() == position.Turn() {
			eval += 100
		}
	}
	return eval
}

func evalBaseWeights(board *chess.Board, position *chess.Position) int {
	eval := 0
	squareMap := board.SquareMap()
	for _, piece := range squareMap {
		if piece.Color() == position.Turn() {
			eval += weights[piece.Type()]
		} else {
			eval -= weights[piece.Type()]
		}
	}
	return eval
}

func evalCanCastle(board *chess.Board, position *chess.Position) int {
	eval := 0
	castleRights := position.CastleRights()
	canCastleKS := castleRights.CanCastle(position.Turn(), chess.KingSide)
	canCastleQS := castleRights.CanCastle(position.Turn(), chess.QueenSide)
	if canCastleQS || canCastleKS {
		eval += 200
	} else {
		eval -= 200
	}

	return eval
}

func evalMobility(board *chess.Board, position *chess.Position) int {
	return len(position.ValidMoves()) * 7
}

func Score(position *chess.Position) int {
	board := position.Board()
	eval := 0

	eval += evalCenterSquares(board, position)
	eval += evalMobility(board, position)
	eval += evalKingSafety(board, position)
	eval += evalBaseWeights(board, position)
	eval += evalCanCastle(board, position)

	return eval + 1
}

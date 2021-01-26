package matechef

import (
	"github.com/notnil/chess"
)

type Context struct {
	alphaPos   *Node
	alphaScore int
	betaPos    *Node
	betaScore  int
}

//finds max of a node given context
func (node *Node) Max(ctx Context) (*Node, int) {
	if node.isTerminal() {
		node.Score()
		return node, node.score
	}
	if node.Position.Status() == chess.Checkmate {
		return node, -999999
	}

	for _, child := range node.children {
		pos, score := child.Min(ctx)

		if score >= ctx.betaScore {
			return ctx.betaPos, ctx.betaScore // fail hard alpha-cutoff
		}
		if score > ctx.alphaScore {
			ctx.alphaPos, ctx.alphaScore = pos, score // beta acts like min in MiniMax
		}
	}
	return ctx.alphaPos, ctx.alphaScore
}

//finds min of a node given context
func (node *Node) Min(ctx Context) (*Node, int) {
	if node.isTerminal() {
		node.Score()
		return node, -node.score
	}
	if node.Position.Status() == chess.Checkmate {
		return node, 999999
	}

	for _, child := range node.children {
		pos, score := child.Max(ctx)
		if score <= ctx.alphaScore {
			return ctx.alphaPos, ctx.alphaScore // fail hard alpha-cutoff
		}
		if score < ctx.betaScore {
			ctx.betaPos, ctx.betaScore = pos, score // beta acts like min in MiniMax
		}
	}
	return ctx.betaPos, ctx.betaScore
}

func (node *Node) Search() *chess.Move {
	ctx := Context{
		alphaPos:   node,
		alphaScore: -999999999,
		betaPos:    node,
		betaScore:  99999999,
	}

	node.GenerateToDepth(4)

	found, _ := node.Max(ctx)
	move := found.ConstructMove()

	return move
}

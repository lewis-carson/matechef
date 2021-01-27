package matechef

import (
	"fmt"

	"github.com/notnil/chess"
)

type Context struct {
	alphaPos   *Node
	alphaScore int
	betaPos    *Node
	betaScore  int
}

type Settings struct {
	MinDepth int
}

func (node *Node) isQuiet() bool {
	return true
}

//finds max of a node given context
func (node *Node) max(ctx Context) (*Node, int) {
	if node.Position.Status() == chess.Checkmate {
		return node, -999999999
	}

	if node.isTerminal() {
		node.Score()
		return node, node.score
	}

	for _, child := range node.children {
		pos, score := child.min(ctx)

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
func (node *Node) min(ctx Context) (*Node, int) {
	if node.Position.Status() == chess.Checkmate {
		return node, 99999999
	}
	if node.isTerminal() {
		node.Score()
		return node, -node.score
	}

	for _, child := range node.children {
		pos, score := child.max(ctx)
		if score <= ctx.alphaScore {
			return ctx.alphaPos, ctx.alphaScore // fail hard alpha-cutoff
		}
		if score < ctx.betaScore {
			ctx.betaPos, ctx.betaScore = pos, score // beta acts like min in MiniMax
		}
	}
	return ctx.betaPos, ctx.betaScore
}

func (node *Node) Search(settings Settings) (*chess.Move, int) {
	ctx := Context{
		alphaPos:   node,
		alphaScore: -999999999998,
		betaPos:    node,
		betaScore:  99999999998,
	}

	node.GenerateToDepth(settings.MinDepth)

	found, score := node.max(ctx)
	fmt.Println(score)
	move := found.ConstructMove()

	return move, score
}

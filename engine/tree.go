package matechef

import (
	"github.com/notnil/chess"
)

type Node struct {
	Position *chess.Position
	parent   *Node
	children []*Node
	move     *chess.Move
	score    int
	isScored bool
}

func reverse(s []*chess.Move) []*chess.Move {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func (node *Node) traceToRoot(line []*chess.Move) []*chess.Move {
	if node.parent != nil {
		line = append(line, node.move)
		line = node.parent.traceToRoot(line)
	}
	return line
}

func (node *Node) ConstructMove() *chess.Move {
	var emptyLine []*chess.Move
	line := reverse(node.traceToRoot(emptyLine))
	return line[0]
}

func (node *Node) isTerminal() bool {
	return len(node.children) == 0
}

func (node *Node) Score() int {
	if !node.isScored {
		node.score = Score(node.Position)
		node.isScored = true
	}
	return node.score
}

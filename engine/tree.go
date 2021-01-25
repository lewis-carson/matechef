package matechef

import "github.com/notnil/chess"

type Node struct {
	Position *chess.Position
	parent   *Node
	children []*Node
	move     *chess.Move
}

func (node *Node) isTerminal() bool {
	return len(node.children) == 0
}

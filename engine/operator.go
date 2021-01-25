package matechef

import (
	"fmt"

	"github.com/notnil/chess"
)

type Context struct {
	depth    int
	alphaPos *Node
	betaPos  *Node
}

//finds max of a node given context
func (node *Node) Max(context Context) *Node {
	return node
}

//finds min of a node given context
func (node *Node) Min(context Context) *Node {
	return node
}

func (node *Node) Search() *chess.Move {
	node.GenerateToDepth(5)
	for _, i := range node.children {
		fmt.Println(i.move)
	}
	return node.Position.ValidMoves()[0]
}

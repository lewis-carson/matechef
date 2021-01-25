package matechef

//adds a single layer of leaves to a node
func (node *Node) Deepen() {
	for _, validMove := range node.Position.ValidMoves() {
		newPos := node.Position.Copy().Update(validMove)
		//wrap position in a node
		newNode := Node{
			parent:   node,
			Position: newPos,
			move:     validMove,
		}
		//append a reference of new node to children
		node.children = append(node.children, &newNode)
	}
}

//add leaves to reference of node until it hits a specified depth
func (node *Node) GenerateToDepth(depth int) {
	//return after bottoming out
	if depth == 0 {
		return
	}
	//add a new layer to node
	node.Deepen()

	//recursively add more children to generated children
	for _, child := range node.children {
		child.GenerateToDepth(depth - 1)
	}
}

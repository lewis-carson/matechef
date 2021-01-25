# mateChef

mateChef is the spiritual successor to another chess engine I made, Kimchi. It aims to be faster and more readable/maintainable. 

It's split into three parts - the node, the generator and the operator.

These are the "brains" of the engine.
- Node - this is the datastructure which holds the relationship between a set of nodes. It holds a variety of attributes.
	- Position
	- Memoized evaluation
	- All children possible to reach within one ply
- Generator - this expands the node, adding more positions and leaf nodes. It works in place, making changes to the node.
- Operator - this is the set of operations which can be applied to the node. The operator walks through the node, making selections and decisions based on an arbitrary set of rules. In a traditional engine, this is the search function.

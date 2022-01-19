package syntax_tree

import "fmt"

type Node struct {
	Operator   string
	Key        string
	Left       *Node
	Right      *Node
	InBrackets bool
}

func (n Node) IsEmpty() bool {
	return n.Operator == "" && n.Key == ""
}

func (n Node) IsOperator() bool {
	return n.Operator != ""
}

func (n *Node) pushLeft(op string) *Node {
	return &Node{
		Operator: op,
		Left:     n,
	}
}

func (n *Node) pushRight(op string) {
	n.Right = &Node{
		Operator: op,
		Left:     n.Right,
	}
}

func (n *Node) pushNot(op string) {
	n.Right = &Node{
		Operator: op,
		Right:    n.Right,
	}
}

func (n Node) print() string {
	if !n.IsOperator() {
		return n.Key
	}

	var left, right string
	if n.Left != nil {
		left = n.Left.print()
	}
	if n.Right != nil {
		right = n.Right.print()
	}

	if n.Operator == OperatorNot {
		return fmt.Sprintf("%s %s", n.Operator, right)
	}
	if n.InBrackets {
		return fmt.Sprintf("(%s %s %s)", left, n.Operator, right)
	}

	return fmt.Sprintf("%s %s %s", left, n.Operator, right)
}

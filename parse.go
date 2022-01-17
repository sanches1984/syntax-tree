package syntax_tree

import "errors"

func parse(expression []string) (*Node, int, error) {
	var cur *Node
	n := &Node{}
	for i := 0; i < len(expression); i++ {
		switch expression[i] {
		case OperatorLBracket:
			x, v, err := parse(expression[i+1:])
			if err != nil {
				return nil, 0, err
			}
			x.strong = true
			i += v + 1
			if n.IsEmpty() {
				n = x
			} else if cur != nil {
				cur.Right = x
			} else if n.Right == nil {
				n.Right = x
			} else {
				return nil, 0, errors.New("tree building error")
			}
		case OperatorRBracker:
			return n, i, nil
		case OperatorNot:
			n.pushNot(expression[i])
			cur = n.Right
		case OperatorAnd:
			if !n.strong {
				n.pushRight(expression[i])
				cur = n.Right
			} else {
				n = n.pushLeft(expression[i])
				cur = n
			}
		case OperatorOr:
			n = n.pushLeft(expression[i])
			cur = n
		default:
			if cur != nil {
				if cur.IsOperator() {
					cur.Right = &Node{Key: expression[i]}
				} else {
					return nil, 0, errors.New("tree building error")
				}
			} else {
				n.Key = expression[i]
			}
		}
	}

	return n, len(expression), nil
}

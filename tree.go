package syntax_tree

const (
	OperatorAnd      = "and"
	OperatorOr       = "or"
	OperatorNot      = "not"
	OperatorLBracket = "("
	OperatorRBracker = ")"
)

type Tree struct {
	Root *Node
}

func Parse(expression []string) (*Tree, error) {
	root, _, err := parse(expression)
	if err != nil {
		return nil, err
	}
	return &Tree{Root: root}, nil
}

func (t Tree) Expression() string {
	if t.Root != nil {
		return t.Root.print()
	}

	return ""
}

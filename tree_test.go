package syntax_tree

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type testCase struct {
	values     []string
	tree       *Tree
	expression string
}

func TestParse(t *testing.T) {
	for _, c := range getTestCases() {
		actual, err := Parse(c.values)

		require.NoError(t, err)
		require.Equal(t, c.tree, actual)
	}
}

func TestExpression(t *testing.T) {
	for _, c := range getTestCases() {
		actual := c.tree.Expression()

		require.Equal(t, c.expression, actual)
	}
}

func getTestCases() []testCase {
	return []testCase{
		{
			values:     []string{"(", "(", "A", "or", "B", ")", "and", "(", "C", "or", "D", ")", ")", "or", "not", "(", "E", "or", "F", "and", "G", ")"},
			expression: "(((A or B) and (C or D)) or not (E or (F and G)))",
			tree: &Tree{
				Root: &Node{
					Operator: OperatorOr,
					Left: &Node{
						Operator: OperatorAnd,
						Left: &Node{
							Operator: OperatorOr,
							Left:     &Node{Key: "A"},
							Right:    &Node{Key: "B"},
							strong:   true,
						},
						Right: &Node{
							Operator: OperatorOr,
							Left:     &Node{Key: "C"},
							Right:    &Node{Key: "D"},
							strong:   true,
						},
						strong: true,
					},
					Right: &Node{
						Operator: OperatorNot,
						Right: &Node{
							Operator: OperatorOr,
							Left:     &Node{Key: "E"},
							Right: &Node{
								Operator: OperatorAnd,
								Left:     &Node{Key: "F"},
								Right:    &Node{Key: "G"},
							},
							strong: true,
						},
					},
				},
			},
		},
	}
}

package main

import (
	"fmt"
	syntax_tree "github.com/sanches1984/syntax-tree"
)

func main() {
	var test = []string{"(", "(", "A", "or", "B", ")", "and", "(", "C", "or", "D", ")", ")", "or", "not", "(", "E", "or", "F", "and", "G", ")"}

	tree, err := syntax_tree.Parse(test)
	if err != nil {
		panic(err)
	}
	fmt.Println(tree.Expression())
}

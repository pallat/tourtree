package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value

	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	var sum1, sum2 int

	n := strings.Count(t1.String(), "(")

	for i := 0; i < n; i++ {
		sum1 += <-ch1
	}

	n = strings.Count(t1.String(), "(")
	for i := 0; i < n; i++ {
		sum2 += <-ch2
	}

	return sum1 == sum2
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(2), tree.New(2)))
	fmt.Println(Same(tree.New(3), tree.New(2)))
	fmt.Println(Same(tree.New(3), tree.New(3)))
}

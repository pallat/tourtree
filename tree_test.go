package main

import (
	"fmt"
	"testing"

	"golang.org/x/tour/tree"
)

func TestWalk(t *testing.T) {
	ch := make(chan int)
	te := tree.New(1)

	go Walk(te, ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

}

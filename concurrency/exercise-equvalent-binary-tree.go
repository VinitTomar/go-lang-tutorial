package concurrency

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func walk(t *tree.Tree, ch chan int) {
	walkRecursion(t, ch)
	close(ch)
}

func walkRecursion(t *tree.Tree, ch chan int) {

	if t != nil {
		walkRecursion(t.Left, ch)
		ch <- t.Value
		walkRecursion(t.Right, ch)
	}

}

func same(t1, t2 *tree.Tree) bool {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go walk(t1, ch1)
	go walk(t2, ch2)

	for {
		n1, ok1 := <-ch1
		n2, ok2 := <-ch2

		if n1 != n2 || ok1 != ok2 {
			return false
		}

		if !ok1 {
			break
		}
	}

	return true

}

// ShowEqBinTree ...
func ShowEqBinTree() {

	println("======================== Binary tree exercise =====================")

	t1 := tree.New(2)
	fmt.Println("T1: ", t1)
	t2 := tree.New(2)
	fmt.Println("T1: ", t2)

	fmt.Println("Same? ", same(t1, t2))

}

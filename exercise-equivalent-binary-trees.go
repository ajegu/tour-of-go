package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {

	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		ch <- t.Value

		if t.Left != nil {
			walk(t.Left)
		}

		if t.Right != nil {
			walk(t.Right)
		}
	}
	walk(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	ch1 := make(chan int)
	go Walk(t1, ch1)

	for v1 := range ch1 {
		flag := false

		ch2 := make(chan int)
		go Walk(t2, ch2)

		for v2 := range ch2 {
			if v1 == v2 {
				flag = true
				break
			}
		}

		if flag == false {
			return false
		}
	}

	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	for v := range ch {
		fmt.Printf("%T(%v)\n", v, v)
	}

	result := Same(tree.New(1), tree.New(1))
	fmt.Println("Comparaison des arbres binaires t1 et t2 : ", result)
}

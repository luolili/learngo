package main

import "fmt"
import ".."

func main() {
	var root tree.Node
	fmt.Println(root)
	root.Left = &tree.Node{}
	root.Left.Right = tree.CreateNode(5)
	root.Right = &tree.Node{Value: 3}
	root.Right.Print()
}

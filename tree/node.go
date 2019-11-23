package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

//工厂函数
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}
func main() {
	var root treeNode
	fmt.Println(root)
	root.left = &treeNode{} //没有构造方法
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)

	nodes := []treeNode{
		{value: 2},
		{},
		{6, nil, &root}, //必须加 ，
	}
	fmt.Println(nodes)
}

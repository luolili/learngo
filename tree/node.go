package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

//给struct添加方法 (node treeNode) 方法所属 的对象
func (node Node) Print() {
	fmt.Println(node.Value)
}

//工厂函数
func CreateNode(value int) *Node {
	return &Node{Value: value}
}

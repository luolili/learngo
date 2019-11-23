package main

import "fmt"

/**
数组 是 值类型
*/
func main() {

	var arr1 [5]int
	arr2 := [3]int{1, 2, 9}
	arr3 := [...]int{1, 2, 9, 0}
	fmt.Println(arr1, arr2, arr3)
	fmt.Println("--", len(arr2))

	for i, v := range arr2 {
		fmt.Println(i, v)
	}
	for _, v := range arr2 {
		fmt.Println(v)
	}
}

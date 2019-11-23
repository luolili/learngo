package main

import "fmt"

/**
slice 是数组的视图：view
*/
//参数是slice
func updateSlice(s []int) {
	s[0] = 100

}

func main() {

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6]", arr[2:6]) //不包含6
	fmt.Println("arr[:6]", arr[:6])
	//fmt.Println("arr[2:]",arr[2:])
	fmt.Println("arr[:]", arr[:])
	s1 := arr[2:]
	fmt.Println("AFTER SLICE")
	updateSlice(s1)
	fmt.Println("s1:", s1)   //s1: [100 3 4 5 6 7]
	fmt.Println("arr:", arr) //s1: [100 3 4 5 6 7]
}

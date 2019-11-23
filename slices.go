package main

import "fmt"

/**
slice 是数组的视图：view,LOOK
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

	fmt.Println("ext:slice")
	arr2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s3 := arr2[2:6]
	s4 := s3[3:5]
	//slice:ptr+len+cap
	fmt.Println("cap:", cap(s3)) //2-6
	fmt.Println("len:", len(s3))
	fmt.Println("s2:", s4)          //可以往后扩展，不可往前扩展
	fmt.Println("s4 cap:", cap(s4)) //2-6
	fmt.Println("len:", len(s4))
}

package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("value:%v,len:%d,cap:%d\n", s, len(s), cap(s))

}
func main() {
	//定义
	var s []int
	fmt.Println(s) //[]
	//fmt.Println(s==nil)//true
	for i := 0; i < 30; i++ {
		//没有值，也不会崩溃
		printSlice(s) //*2 自增
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
	//其他定义
	s1 := []int{1, 2}
	printSlice(s1)
	s2 := make([]int, 16)
	printSlice(s2)
	s3 := make([]int, 10, 22)
	printSlice(s3)
}

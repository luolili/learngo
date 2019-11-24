package main

import (
	"fmt"
)

func tryDefer() {
	defer fmt.Println(1) //函数退出之时 执行
	defer fmt.Println(2)
	fmt.Println(3)

}

func tryRecover() {

	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("er:", err)
		} else {
			panic(fmt.Sprintf("i do not know what to do,%v", r))

		}
	}()
	//panic(errors.New("it is an error"))
	//b := 0
	//a := 3 / b //er: runtime error: integer divide by zero
	//fmt.Println(a)
	panic(11)
}
func main() {
	//tryDefer()
	tryRecover()
}

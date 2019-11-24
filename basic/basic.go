package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Println(a, s)
	fmt.Printf("%d %q\n", a, s)
}

func variableInitValue() {
	var a int = 2
	var s string = "sw"
	fmt.Println(a, s)
}

func variableTypeDeduction() {

	var a, b, c = 1, "we", 3
	fmt.Println(a, b, c)
	//在方法外面不可这样定义var
	d := 2
	fmt.Println(d)

}

func euler() {
	/*c:=3+4i
	fmt.Println(cmplx.Abs(c))*/
	//cmplx :数学里面的复数，浮点数的标准 就说明了他的不准确，任何语言遵循了这个标准都是imprecise
	c := cmplx.Pow(math.E, 1i*math.Pi) + 1
	fmt.Println(c)
}

func triangle() {
	var a, b = 3, 4
	fmt.Println(calcTriangle(a, b))
}
func calcTriangle(a, b int) int {
	var c int
	//只有强制的类型转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c

}
func consts() {
	const filename = "ad"
	const a, b = 3, 4 //没有指定类型，他就是个文本值
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c, filename)

}

//也是 const
func enums01() {
	const (
		//必须赋值
		java   = 0
		python = 1
		golang = 2 //不能用go,
	)
	fmt.Println(java, python, golang)

}
func enums02() {
	const (
		//必须赋值,自增1
		java = iota
		_
		python
		golang
	)
	fmt.Println(java, python, golang)

}
func enums03() {
	const (
		//必须赋值,自增 制定的值
		java = 1 << (iota * 10)
		python
		golang
	)
	fmt.Println(java, python, golang)

}

//包内的var
var name = "00"

/**
类型：(u) int,int8,int16,int32,int64
float32,float64
const,enum,rune
*/
func main() {
	fmt.Println("halo world")
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	fmt.Println(name)
	euler()
	triangle()
	consts()
	enums01()
	enums02()
	enums03()

}

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

//包内的var
var name = "00"

func main() {
	fmt.Println("halo world")
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	fmt.Println(name)
	euler()
	triangle()
}
func euler() {
	/*c:=3+4i
	fmt.Println(cmplx.Abs(c))*/
	c := cmplx.Pow(math.E, 1i*math.Pi) + 1
	fmt.Println(c)

}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)

}

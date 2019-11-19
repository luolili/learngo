package main

import "fmt"

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
func main() {
	fmt.Println("halo world")
	variableZeroValue()
	variableInitValue()
}

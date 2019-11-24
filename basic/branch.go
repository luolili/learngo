package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	re := ""
	switch {
	case score < 60:
		re = "F"

	case score < 90:
		re = "B"
	case score > 100:
		re = "A"

	default:
		panic(fmt.Errorf("wrong score %d", score))

	}
	return re

}

//不能有括号
func main() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename) //返回了文本和错误信息
	//open abc.txt: The system cannot find the file specified.没有该文件的报错
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents) //contents 是byte[]
	}
	fmt.Println(grade(90))
}

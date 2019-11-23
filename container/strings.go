package main

import (
	"fmt"
)

func main() {
	s := "我是什么?" //UTF-8
	fmt.Println(len(s))
	fmt.Printf("%s", []byte(s))
	fmt.Printf("%X", []byte(s))

	for _, v := range []byte(s) {
		fmt.Printf("%X  ", v)
	}
	fmt.Println()
	//fmt.Println("rune count:",utf8.RuneCountInString(s))
	//bytes:=[]byte(s)
	/* for len(bytes)>0{
	 	ch,size:=utf8.DecodeRune(bytes)
	 	bytes=bytes[size:]
		 fmt.Println(ch,size)
	 }*/

	for i, ch := range []rune(s) {
		fmt.Printf("(%d,%c)", i, ch)
	}

}

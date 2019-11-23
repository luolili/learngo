package main

import "fmt"
import "./mock"
import "./real"

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://studygolang.com")

}
func main() {
	var r Retriever
	r = mock.Retriever{"this is fake"}
	fmt.Printf("%T,%v\n", r, r) //类型和值
	r = real.Retriever{}
	//fmt.Println(download(r))
}

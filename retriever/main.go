package main

import "fmt"
import "./mock"
import "./real"

//接口
type Retriever interface {
	Get(url string) string
}
type Poster interface {
	Post(url string, form map[string]string) string
}

//接口 的组合

type RetrieverPoster interface {
	Retriever
	Poster
}

func post(poster Poster) {
	poster.Post("https://studygolang.com",
		map[string]string{
			"name": "mee",
			"age":  "11",
		})

}

const url = "https://studygolang.com"

func session(s RetrieverPoster) string {

	s.Post(url, map[string]string{
		"contents": "ff",
	})
	return s.Get(url)
}

func download(r Retriever) string {
	return r.Get("https://studygolang.com")

}
func main() {
	var r Retriever
	r = mock.Retriever{"this is fake"}
	fmt.Printf("%T,%v\n", r, r) //类型和值
	r = real.Retriever{}
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("-mock-", v.Contents)
	case real.Retriever:
		fmt.Println("-real-", v.UserAgent)

	}
	//fmt.Println(download(r))

	receiver := mock.Retriever{"this is test"}

	fmt.Println(session(receiver))
}

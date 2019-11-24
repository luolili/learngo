package main

import (
	"fmt"
	"time"
)

func main() {

	chanDemo()
}

func chanDemo() {
	//var c chan int//chanel里面是int
	c := make(chan int)
	/*go func() {
		//收数据
		n:=<-c
		fmt.Println(n)
	}()*/
	go worker(0, c)
	//放数据
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)

}
func worker(id int, c chan int) {
	for {
		fmt.Printf("workder %d,received %d", id, <-c)
	}

}

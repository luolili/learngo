package main

import (
	"fmt"
	"time"
)

func main() {
	//var a [10]int
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println("go routine from", i)
			/*for{
				//a[i]++
				//runtime.Gosched()//交出控制权
			}*/
		}(i)

	}
	time.Sleep(time.Minute) //1毫秒
}

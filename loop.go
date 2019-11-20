package main

import (
	"fmt"
	"strconv"
)

func convertToBin(n int) string {

	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

//死循环，用处多
func forever() {
	for {
		fmt.Println(11)
	}

}
func main() {

	fmt.Println(convertToBin(5))
	forever()
}

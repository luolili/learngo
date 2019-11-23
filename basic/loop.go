package basic

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

//死循环，用处多，没有while
func forever() {
	for {
		fmt.Println(11)
	}

}

func sum(nums ...int) int {

	s := 0
	for i := range nums {
		s += nums[i]
	}
	return s
}
func main() {

	fmt.Println(convertToBin(5))
	//forever()
	// 可变 参数
	fmt.Println(sum(1, 2, 3))
}

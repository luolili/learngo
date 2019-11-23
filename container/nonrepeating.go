package main

import "fmt"

func noRepeat(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastI = i
	}
	return maxLength
}
func noRepeat01(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastI = i
	}
	return maxLength
}
func main() {

	fmt.Println(noRepeat("adcdwec"))
	fmt.Println(noRepeat("abcde"))
	fmt.Println(noRepeat("  "))
	fmt.Println(noRepeat("这是一个的")) //不支持中文
	fmt.Println(noRepeat01("这是一个的"))
}

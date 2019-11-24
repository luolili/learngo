package main

import (
	"fmt"
	"regexp"
)

const text = "mu mail is xx@qq.com@de"

func main() {
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	match := re.FindString(text)
	fmt.Printf(match)
}

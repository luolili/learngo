package main

import "fmt"

/**
map
*/
func main() {
	m := map[string]string{
		"name": "wee",
		"age":  "11",
		"addr": "sw",
	}
	fmt.Println(m) //逆 打印
	m2 := make(map[string]int)
	var m3 map[string]int
	fmt.Println(m2, m3)
	fmt.Println(m2 == nil) //false
	fmt.Println(m3 == nil) //true

	fmt.Println("traversing map")
	for k, v := range m {
		fmt.Println(k, v) //每次打印的无序的
	}
	for k, _ := range m {
		fmt.Println(k)
	}
	fmt.Println("getting value")
	name := m["name"]
	fmt.Println("name value:", name)
	nam := m["nam"]
	fmt.Printf("nam value:%q\n", nam)
	//addr,ok:=m["nam"]
	//fmt.Printf("nam value:\n",addr,ok)
	//fmt.Printf("nam value:\n",addr,ok==false)

	delete(m, "name")
	fmt.Println("getting value:", m)
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic("Error")
		return
	}
	if resp.StatusCode == http.StatusOK {
		all, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("%s\n",all)

		printCityList(all)

	}

}
func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("City:%s,url:%s \n ", m[2], m[1])
		/*for _,subMatch:= range m{
			//fmt.Printf("%s  ",subMatch)
		}*/
		fmt.Println()
	}
	fmt.Printf("%d\n", len(matches))

}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("https://www.zhenai.com/zhenghun")
	if err!= nil{
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}
	//fmt.Printf("%T", resp.Body)
	all, err := ioutil.ReadAll(resp.Body)		// 返回网页源码
	if err != nil{
		panic(err)
	}
	//fmt.Printf("%s ", all)
	printCityList(all)
}

func printCityList(contents []byte) {



	re := regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)".data-v-5e16505f[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range  matches{
		fmt.Printf("City: %s, URL: %s\n", m[2], m[1])


	}
	fmt.Println(len(matches))


}

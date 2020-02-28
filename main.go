package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s ", all)




}

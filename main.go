package main

import (
	"crawZhenai/engine"
	"crawZhenai/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/",
		ParserFunc: parser.ParseCityList, // 解释的方法为ParseCityList()
	})

}

package main

import (
	"crawZhenai/engine"
	"crawZhenai/scheduler"
	"crawZhenai/zhenai/parser"
	"fmt"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkCount: 100,
	}
	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/",
	//	ParserFunc: parser.ParseCityList, // 解释的方法为ParseCityList()
	//})
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/guangan",
		ParserFunc: parser.ParseCity, // 解释的方法为ParseCityList()
	})

	fmt.Println("所有数据获取完毕！")

}

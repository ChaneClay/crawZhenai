package main

import (
	"crawZhenai/engine"
	"crawZhenai/scheduler"
	"crawZhenai/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkCount: 100,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/",
		ParserFunc: parser.ParseCityList, // 解释的方法为ParseCityList()
	})

}

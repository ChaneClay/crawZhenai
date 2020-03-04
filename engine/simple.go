package engine

import (
	"crawZhenai/fetcher"
	"fmt"
	"log"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request //seeds 为[]engine.Request类型

	for _, r := range seeds {
		requests = append(requests, r) //	把r存储到requests
	}

	for len(requests) > 0 { //开始只有一个请求				//第一轮把所有城市名和url拿到
		r := requests[0]        // 把第一个拿出来				//第二轮进入第一个城市里获取所有用户名和其对应主页的url
		requests = requests[1:] // 元素左移					//第三轮进入第二个城市里获取所有用户名和其对应主页的url
		//第n轮进入第一个城市的第一个用户的主页获取资料

		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		//fmt.Println("-----------------------------------")
		//fmt.Println(parseResult.Requests[0].Url, parseResult.Requests[1].Url, parseResult.Requests[2].Url)
		//fmt.Println("-----------------------------------")

		requests = append(requests, parseResult.Requests...) //parseResult.Requests的元素被打散一个个append进requests
		// 把所有城市的链接存放到requsts, 打印城市名字
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
		fmt.Println("所有数据获取完毕！")
	}

}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s ", r.Url)
	body, err := fetcher.Fetch(r.Url) //传递url返回网页源码

	if err != nil {
		log.Printf("Fetcher: err fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil // 调用ParseCityList方法, 获取url和城市名字
}

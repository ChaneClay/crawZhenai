package engine

import (
	"crawZhenai/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request //seeds 为[]engine.Request类型

	for _, r := range seeds {
		requests = append(requests, r) //	把r存储到requests
	}

	for len(requests) > 0 { //开始只有一个请求				//第一轮把所有城市名和url拿到
		r := requests[0]        // 把第一个拿出来				//第二轮进入第一个城市里获取所有用户名和其对应主页的url
		requests = requests[1:] // 元素左移					//第三轮进入第二个城市里获取所有用户名和其对应主页的url
		//第n轮进入第一个城市的第一个用户的主页获取资料
		log.Printf("Fetching %s ", r.Url) //第n+1轮进入第一个城市的第二个用户的主页获取资料
		body, err := fetcher.Fetch(r.Url) //传递url返回网页源码

		if err != nil {
			log.Printf("Fetcher: err fetching url %s: %v", r.Url, err)
			continue
		}
		parseResult := r.ParserFunc(body) // 调用ParseCityList方法, 获取url和城市名字

		requests = append(requests, parseResult.Requests...) //parseResult.Requests的元素被打散一个个append进requests
		// 把所有城市的链接存放到requsts, 打印城市名字
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}

}

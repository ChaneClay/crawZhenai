package engine

import (
	"crawZhenai/fetcher"
	"fmt"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request //seeds 为[]engine.Request类型

	for _, r := range seeds {
		requests = append(requests, r) //	把r存储到requests
	}

	for len(requests) > 0 {
		r := requests[0]        // 把第一个拿出来
		requests = requests[1:] // 元素左移

		log.Printf("Fetching %s ", r.Url)
		body, err := fetcher.Fetch(r.Url)

		if err != nil {
			log.Printf("Fetcher: err fetching url %s: %v", r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(body) // 调用ParseCityList方法

		if parseResult.Items == nil {
			fmt.Println("parseResult.Items is nil.")
		}

		requests = append(requests, parseResult.Requests...) //parseResult.Requests的元素被打散一个个append进requests
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}

}

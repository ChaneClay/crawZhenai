package parser

import (
	"crawZhenai/engine"
	"regexp"
)

const cityListRe = `href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)".data-v-5e16505f[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{} // 初始化变量

	for _, m := range matches { // 遍历获得的city, url, 让后放到ParseResult结构体存储
		result.Items = append(result.Items, "City: "+string(m[2])) // 城市名字
		result.Requests = append(result.Requests, engine.Request{  // 对应的城市url
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})

	}

	return result //返回item、requests。requests里面存储url和解释的方法ParserFunc。ParserFunc 的返回值又是item、requests。

}

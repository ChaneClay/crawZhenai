package parser

import (
	"crawZhenai/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{} // 初始化变量
	for _, m := range matches {    // 遍历获得的city, url, 让后放到ParseResult结构体存储
		name := string(m[2])
		result.Items = append(result.Items, "user: "+name)        // 城市名字
		result.Requests = append(result.Requests, engine.Request{ // 对应的城市url
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})
	}
	return result //返回item、requests。requests里面存储url和解释的方法ParserFunc。ParserFunc 的返回值又是item、requests。
}

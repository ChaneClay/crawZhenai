package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult // ParseResult 为返回值类型
}

type ParseResult struct { // 存储解释后的信息
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}

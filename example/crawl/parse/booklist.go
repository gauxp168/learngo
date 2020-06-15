package parse

import (
	"regexp"
	"simqo.com/mygospace/learngo/example/crawl/engine"
)

const BokkList = `<a href="([^"]+)" title="([^"]+)" `

func ParseBookList(content []byte) engine.ParseResult {
	re := regexp.MustCompile(regexpStr)
	match := re.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	for _, m := range match {
		bookname:= string(m[2])
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Requset{
			Url:string(m[1]),
			ParseFunc: func(bytes []byte) engine.ParseResult {
				return ParseBookDetail(bytes, bookname)
			},
		})
	}
	return result
}
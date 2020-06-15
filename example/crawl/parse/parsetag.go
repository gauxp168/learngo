package parse

import (
	"regexp"
	"simqo.com/mygospace/learngo/example/crawl/engine"
)

const regexpStr = `<a href="([^"]+)" class="tag">([^"]+)</a>`

func ParseContent(content []byte) engine.ParseResult {
	re := regexp.MustCompile(regexpStr)
	match := re.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	for _, m := range match {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Requset{
			Url:"https://book.douban.com/" +string(m[1]),
			ParseFunc:ParseBookList,
		})
	}
	return result
}

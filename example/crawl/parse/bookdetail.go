package parse

import (
	"regexp"
	"simqo.com/mygospace/learngo/example/crawl/engine"
	"simqo.com/mygospace/learngo/example/crawl/model"
	"strconv"
)

var authorRe = regexp.MustCompile(`<span class="pl"> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`)
var publicRe = regexp.MustCompile(`<span class="pl">出版社:</span>([^<]+)<br/>`)
var pageRe = regexp.MustCompile(`<span class="pl">页数:</span>([^<]+)<br/>`)
var priceRe = regexp.MustCompile(`<span class="pl">定价:</span>([^<]+)<br/>`)
var scoreRe = regexp.MustCompile(`<strong class="ll rating_num " property="v:average">([^<]+)</strong>`)
var infoRe = regexp.MustCompile(`<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`)

func ParseBookDetail(content []byte, name string) engine.ParseResult {
	bookdetail := model.Bookdetail{}
	bookdetail.Name = name
	bookdetail.Author = ExtraString(content, authorRe)
	bookdetail.Publicer = ExtraString(content, publicRe)
	page, err := strconv.Atoi(ExtraString(content, pageRe))
	if err == nil {
		bookdetail.Bookpages = page
	}
	bookdetail.Price = ExtraString(content, priceRe)
	bookdetail.Score = ExtraString(content, scoreRe)
	bookdetail.Info = ExtraString(content, infoRe)
	result := engine.ParseResult{
		Items:[]interface{}{bookdetail.String()},
	}
	return result
}

func ExtraString(content []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(content)
	if len(match) > 2 {
		return string(match[1])
	}else {
		return ""
	}
}

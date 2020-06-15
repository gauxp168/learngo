package engine


import (
	"fmt"
	"simqo.com/mygospace/learngo/example/crawl/fetcher"
)

type SimpleEngine struct {

}

func (s SimpleEngine) Run(seeds ...Requset)  {
	var requests []Requset
	for _, e := range seeds {
		requests = append(requests, e)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			fmt.Println("error:", err)
		}
		parseresult := r.ParseFunc(body)
		requests = append(requests, parseresult.Requests...)
		for _, item := range parseresult.Items {
			fmt.Printf("get item:%s\n", item)
		}
	}
}

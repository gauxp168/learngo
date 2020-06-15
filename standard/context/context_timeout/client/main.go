package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"simqo.com/old_boy_5_Golang/day17课上代码和笔记/项目/项目资料/vendor包/google.golang.org/appengine/log"
	"sync"
	"time"
)

type respData struct {
	resp *http.Response
	err error
}

func doCall(ctx context.Context)  {
	transport := http.Transport{
		// 请求频繁可定义全局的client对象并启用长链接
		// 请求不频繁使用短链接
		DisableKeepAlives:true,
	}

	client := http.Client{
		Transport:&transport,
	}

	respChan := make(chan *respData, 1)
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080", nil)
	if err != nil {
		fmt.Printf("new request failed, error:%v\n", err)
		return
	}
	req = req.WithContext(ctx)	// 使用带超时的ctx创建一个新的client request
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go func() {
		resp, err := client.Do(req)
		fmt.Printf("client do resp:%v, err:%v\n", resp, err)
		rd := &respData{
			resp:resp,
			err:err,
		}
		respChan<-rd
		wg.Done()
	}()

	select {
	case <-ctx.Done():
		// transport.CancelRequest(req)
		fmt.Println("call api timeout")
	case result := <-respChan:
		fmt.Println("call server api success")
		if result.err != nil {
			fmt.Printf("call server api failed, error:%v\n", err)
			return
		}
		defer result.resp.Body.Close()
		data, _ := ioutil.ReadAll(result.resp.Body)
		fmt.Printf("resp:%v\n", string(data))
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	doCall(ctx)
}














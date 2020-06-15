package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)
var ratelimit = time.Tick(1000*time.Millisecond)
func Fetch(weburl string) ([]byte, error) {
	<-ratelimit
	proxy := func(_ *http.Request) (*url.URL, error){
		return url.Parse("http://127.0.0.1:1087")
	}
	transport := &http.Transport{Proxy:proxy}
	client := &http.Client{Transport:transport}
	request, err := http.NewRequest("GET", weburl, nil)
	if err != nil {
		return nil, fmt.Errorf("ERROR:got url:%s", err)
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("error status code:%d\n", resp.StatusCode)
	}
	reader := bufio.NewReader(resp.Body)
	e := DeterminEncoding(reader)
	newReader := transform.NewReader(reader, e.NewDecoder())
	return ioutil.ReadAll(newReader)
}
func WebFetch(url string) ([]byte, error) {
	<-ratelimit
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("ERROR:got url:%s", err)
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("error status code:%d\n", resp.StatusCode)
	}
	reader := bufio.NewReader(resp.Body)
	e := DeterminEncoding(reader)
	newReader := transform.NewReader(reader, e.NewDecoder())
	return ioutil.ReadAll(newReader)
}

func DeterminEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("fetch error:%v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}


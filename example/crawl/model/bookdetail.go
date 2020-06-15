package model

import "strconv"

type Bookdetail struct {
	Name string
	Author string
	Publicer string
	Bookpages int
	Price string
	Score string
	Info string
}

func (b Bookdetail) String() string {
	return "书名："+b.Name+"作者:"+b.Author+"出版社："+b.Publicer+"书页："+strconv.Itoa(b.Bookpages)+"书价："+b.Price+"评分："+b.Score+"\n简介：\n"+b.Info
}
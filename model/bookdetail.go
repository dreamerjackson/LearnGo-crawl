package model

import "strconv"

type  Bookdetail struct {
	BookName string
	Author string
	Publicer string
	Bookpages int
	Price string
	Score string
	Into  string

}

func (b Bookdetail) String() string{

	return "书籍名字:" + b.BookName+ " 作者 :" + b.Author + " 出版社"+ b.Publicer + " 书籍页数："+ strconv.Itoa(b.Bookpages)+ " 价格："+ b.Price +" 得分"+ b.Score + " \n简介:"+ b.Into
}
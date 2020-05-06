// 结构体的嵌套之一

package main

import (
	"fmt"
)

type Book struct {
	BookName string
	BookPrice int
}

type Stu struct {
	StuName string
	StuAge int
	StuFavBook Book
}

func main() {


	s1 := Stu{
		StuName : "pengw",
		StuAge : 20,
		StuFavBook : Book{
			BookName : "Go语言",
			BookPrice : 20,
		},
	}

	fmt.Println(s1, s1.StuFavBook.BookName)  // {pengw 20 {Go语言 20}} Go语言
}
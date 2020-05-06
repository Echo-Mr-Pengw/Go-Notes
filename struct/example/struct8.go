// 结构体的嵌套之二

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
	StuFavBook *Book
}

func main() {

	b := Book{
		BookName:  "Go",
		BookPrice: 10,
	}

	s := Stu{
		StuName:    "pengw",
		StuAge:     20,
		StuFavBook: &b,
	}

	fmt.Println(s, *s.StuFavBook)  // {pengw 20 0xc00004c420} {Go 10}
}
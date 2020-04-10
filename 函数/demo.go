package main 

import "fmt"


func demo(args...int) {
	fmt.Println(args[0]);
}

func main() {
	demo(1,2,3,4)
}
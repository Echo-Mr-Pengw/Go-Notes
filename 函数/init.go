//每个go文件都又严格init方法
//执行顺序；全局变量 > init() > main()


package main 

import "fmt"

//全局变量
var g = test()

func test() int {
	fmt.Println("全局变量");
	return 1
}

func init() {
	fmt.Println("init");
}

func main() {
	fmt.Println("main", g);
}

//output:
//全局变量
//init
//main 1
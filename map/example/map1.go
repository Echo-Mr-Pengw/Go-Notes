// map的声明及初始化

package main 

import "fmt"

func main() {

	//第一种方法
	//声明
	var map1 map[int]string
	//make 分配内存空间
	map1 = make(map[int]string)

	map1[0] = "刘备"
	map1[1] = "诸葛亮"
	fmt.Println(map1)   // output: map[0:刘备 1:诸葛亮]

	//第二种方法
	map2 := make(map[string]string)
	map2["0"] = "宋江"
	map2["1"] = "武松"
	fmt.Println(map2)  // output: map[0:宋江 1:武松]

	//第三种方法
	map3 := map[string]string{
		"0" : "红楼梦",
		"1" : "西游记",
	}

	fmt.Println(map3)  // output: map[0:红楼梦 1:西游记]
}
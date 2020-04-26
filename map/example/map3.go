// map 是引用类型

package main 

import "fmt"

func test(map1 map[int]int) {
	map1[0] = 100
	fmt.Println(map1)   // output: map[0:100 1:2 2:3]
}

func main() {

	map1 := map[int]int{0 : 1, 1 : 2, 2 : 3}

	test(map1)

	fmt.Println(map1)   // output: map[0:100 1:2 2:3]
}
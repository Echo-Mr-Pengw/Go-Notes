// break 用于for switch select
// break 搭配标签使用

package main

import "fmt"

func main() {
	exitFor:   // 注意：只能放在上面使用，与14行标签配对
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				fmt.Printf("i=%v, j=%v \n", i, j) // output: i=0, j=1
				break exitFor // 直接退出最外层循环
			}
		}
	}
	fmt.Println("exit")
}
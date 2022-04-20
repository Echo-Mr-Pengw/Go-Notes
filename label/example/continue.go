// continue 一般用于for
// continue 搭配标签使用

package main

import "fmt"

func main() {
exitFor: // 注意：只能放在上面使用，与14行标签配对；跟break搭配标签一样
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				fmt.Printf("i=%v, j=%v \n", i, j) // output: i=0, j=1;  i=1; j=1; i=2, j=1
				continue exitFor //退出第二层循环后，在进行for循环，直到不满足条件
			}
		}
	}
		fmt.Println("exit")  // output: exit
}

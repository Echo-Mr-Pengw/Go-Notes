// switch fallthrough: switch穿透
// 如果case后面有fallthrough，会继续执行下一个case语句

package main 

import "fmt"

func main() {

	i := 1

	switch i {
		case 1:
			fmt.Println(1)      //output: 1
			fallthrough
		case 2:
			fmt.Println(2);     //output: 2
		default:
			fmt.Println("no~~~~");
	}
}
//闭包函数：函数+引用

package main 

import "fmt"

//函数test()   返回值 func (int, int) (int, int)
func test() func (int, int) (int, int) {
	i := 1
	j := 2
	return func (x, y int) (int, int) {
		i += x
		j += y
		return i, j
	}

}

func main() {
	t := test()

	fmt.Println(t(1, 2));
	fmt.Println(t(1, 2));
}

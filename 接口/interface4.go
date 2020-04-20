// 接口是引用类型
// 空接口返回nil
 
package main 

import(
	"fmt"
)

type A interface {

}

func main() {
	var a A

	fmt.Println(a)  // output: <nil>
}
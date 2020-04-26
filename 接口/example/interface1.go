// 接口的定义及使用

package main 

import(
	"fmt"
)

type A interface {
	Say()
}

type B struct {

}

func (b B) Say() {
	fmt.Println("Say().....")
}

func main() {
	
	var b B
	var a A = b
	a.Say()
	// output:
	// Say().....
}
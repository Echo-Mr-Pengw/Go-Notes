// 自定义错误  errors.New()和panic
// errors.New("错误信息")  返回一个error类型的错误信息
// painc 内置函数，输出错误信息，并终止程序。


package main 

import(
	"fmt"
	"errors"
)

func errorsNew(name string) (err error) {
	if name == "pengw" {
		return nil
	}

	return errors.New("name error")
}

func paincFunc(name string) {

	err := errorsNew(name)
	if err != nil {
		panic(err)      // output: "name error
	}

	fmt.Println("this name is correct.")
}

func main() {

	//name := "pengw"
	name := "1"
	paincFunc(name)  // 后面的代码不执行

	fmt.Println("next...")
}
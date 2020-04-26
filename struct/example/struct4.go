// 结构体 标签

package main 

import (
	"fmt"
	"encoding/json"
)

type User struct {
	Name string `json:"name"`  // 注意 冒号前后没有空格（`json : name`）
} 

func main() {

	u := User{"张三"}
	json, err := json.Marshal(u)     // 第一个值是字节切片类型，第二个是错误类型
	fmt.Println(string(json), err)   // 字节切片转换成成字符串
}
package main

import (
	"fmt"
	"reflect"
)

type x int

type Cat struct {
	Color 	string
	Age 	x `json:"age"`
}

func main() {

	cat := Cat{
		Color: "白色",
		Age :2,
	}

	tpVal := reflect.ValueOf(cat)

	// 通过ValueOf 获取结构体的字段数
	structFileNum := tpVal.NumField()
	fmt.Println(structFileNum)  // output: 2

	// 	通过TypeOf 获取结构体的字段数
	// fmt.Println(reflect.TypeOf(cat).NumField()) // 2

	tp := reflect.TypeOf(cat)
	secondField := tp.Field(1)
	// 获取字段的类型
	fmt.Println(secondField.Type.Name())   // output: x
	// 获取字段的种类
	fmt.Println(secondField.Type.Kind())   // output: int
	// 获取字段名称
	fmt.Println(secondField.Name)          // output: Age
	// 获取字段tag
	fmt.Println(secondField.Tag.Get("json"))  // output: age； 注意"json"可换成其它对它，但是解析的时候必须为json
}

package main

import (
	"fmt"
	"reflect"
)

type X int

type Student struct {}

func main() {

	// TypeOf
	// Name是类型
	// Kind是种类（可以理解为数据类型）
	var i X
	tyOf := reflect.TypeOf(i)
	fmt.Println(tyOf.Name(), tyOf.Kind())  // output: X  int

	// ValueOf
	var stu Student
	tyVal := reflect.ValueOf(stu)
	fmt.Println(tyVal.Kind()) // 注意：没有直接的Name()方法； output: struct
	fmt.Println(tyVal.Type().Name(), tyVal.Type().Kind())  // output: Student struct
}

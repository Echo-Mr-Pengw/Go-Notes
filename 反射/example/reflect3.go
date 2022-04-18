package main

import (
	"fmt"
	"reflect"
)

func updateVal(i *int) {
	val := reflect.ValueOf(i)
	val.Elem().SetInt(20)  // 注意: Elem()
}

func main() {
	i := 10
	updateVal(&i)
	fmt.Println(i)  // output: 20
}
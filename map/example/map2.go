// map的增删改查及遍历

package main 

import "fmt"

func main() {

	map1 := map[string]string{
		"1" : "C",
		"2" : "Java",
		"3" : "PHP",
		"4" : "Go",
	}

	fmt.Println(map1)   // map[1:C 2:Java 3:PHP 4:Go]

	//添加
	map1["5"] = "HTML"
	fmt.Println(map1)  // map[1:C 2:Java 3:PHP 4:Go 5:HTML]

	//修改
	map1["5"] = "CSS"
	fmt.Println(map1) // map[1:C 2:Java 3:PHP 4:Go 5:CSS]

	//查找
	v, isExist := map1["1"] // isExist:是否存在，v:表示值。isExist=true, v有数值，isExist=false, v无数值，
	fmt.Println(v, isExist) // C true

	//删除 有key则删除 没key不删除且不报错
	delete(map1, "0")
	fmt.Println(map1) // map[2:Java 3:PHP 4:Go 5:CSS]

	//遍历
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	//output
	// 3 PHP
	// 4 Go
	// 5 CSS
	// 1 C
	// 2 Java
}
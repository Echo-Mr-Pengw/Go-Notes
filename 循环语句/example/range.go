// for range

package main

import "fmt"

func main() {

	// range会复制原数据
	a := [2]int{1, 2}
	for i, v := range a {
		if i == 0 {
			a[0] += 10
			a[1] += 10
		}
		//   1   11
		//   2   22
		fmt.Println(v, a[i])
	}

	s := []int{1, 2}
	for i, v := range s {
		if i == 0 {
			s[0] += 10
			s[1] += 10
		}
		//  1  11
		// 12  12
		fmt.Println(v, s[i])
	}
}
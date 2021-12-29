// iota 中断

package main

import "fmt"

func main() {

	const (
		a = 1
		b = iota
		c = 100
		d
		e = iota
		f
	)
	fmt.Println(a, b, c, d, e, f)  // 1 1 100 100 4 5
}
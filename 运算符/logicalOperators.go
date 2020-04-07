//逻辑运算符

package main 

import "fmt"

func main() {

	// &&
	fmt.Println(true && false, true && true, false && true, false && false); //output: false  true  false  false

	// ||
	fmt.Println(true || false, true || true, false || true, false || false); //output: true  true  true  false

	// !
	fmt.Println(!true, !false);  //output: false  true
}
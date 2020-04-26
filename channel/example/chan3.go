// 管道不关闭，使用管道会出现阻塞 deadlock
// select 解决管道阻塞问题 

package main 

import(
	"fmt"
)

func main() {

	ch1 := make(chan int, 10)
	ch2 := make(chan string, 10)

	for i := 0; i < 10; i++ {
		ch1<- i
	}

	for j := 0; j < 10; j++ {
		ch2<- "go" + fmt.Sprintf("%d", j)
	}

	for {
		// 如果第一个管道ch1取不到数据，不会阻塞，会执行下一个case
		select {
			case num1 := <-ch1:
				fmt.Println("从管道ch1中取值：", num1)
			case num2 := <-ch2:
				fmt.Println("从管道ch2中取值：", num2)
			default:
				return
		}
	}
}


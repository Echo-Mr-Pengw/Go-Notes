// 三个索引的使用
// slice[i:j:k]
// i:起始位置，j:结束位置(不包括)，k:该切片容量
// j=k的时候添加元素，底层会生产新的数组
// j<k的时候添加元素，之前的元素会被覆盖
// j>k panic

package main

import "fmt"

func main() {

	s1 := []int{1, 2, 3, 4, 5}

	// j=k
	s2 := s1[0:1:1] // j=k

	fmt.Printf("s2=%v\n", s2)

	// 在s2中添加元素
	s2 = append(s2, 6)  // s2=[1]
	// 可以看出s1中的值没变
	fmt.Printf("s1=%v, s2=%v\n", s1, s2)   // output: s1=[1 2 3 4 5], s2=[1 6]

	// j<k
	s3 := s1[0:1:2]
	s3 = append(s3, 7)
	// 可以看出s1中的值变了
	fmt.Printf("s1=%v, s3=%v\n", s1, s3)   // output: s1=[1 7 3 4 5], s3=[1 7]

	// s4 := s1[0:2:1] 错误
}

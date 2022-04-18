package example

import "testing"

// 单测
// 1. Test+首字母大写的方法名，此方法名可自己指定，非getUserInfo，比如可用：TestInfo
// 2. 方法参数固定值：t *testing.T
// 3. 当前目录执行 go test 执行当前目录下面所有的单测
// 4. 当前目录执行 go test -v 执行当前目录下面所有的单测，并且返回详细信息
// 5. 当前目录执行 go test -v xx_test.go  xx.go， 执行xx_test.go中的单侧，不是当前目前下所有的单测
// 6.  当前目录执行 go test -v -test.run 单侧方法名，执行指定单侧的方法名
func TestGetUserInfo(t *testing.T) {
	name := getUserInfo("testing")
	t.Log(name)
}

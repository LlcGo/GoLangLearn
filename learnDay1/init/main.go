package main

import (
	// 1.导包 可用别名
	// 2.需要只用包内的init函数 _ "GoLearn/init/a"
	// 3.需要里面所有的方法 . "GoLearn/init/a" 不建议使用
	test "GoLearn/init/a"
	lib "GoLearn/init/lib"
)

func main() {
	lib.Lib1Func()
	test.Lib2Func()
}

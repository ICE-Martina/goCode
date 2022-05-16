package main

import (
//	"org.lw/chapter4/lib1"
	_ "org.lw/chapter4/lib1" // 匿名导包, 导入可以不使用， 但是init方法会执行
	// "org.lw/chapter4/lib2"
	// mylib2 "org.lw/chapter4/lib2" // 给包起别名
	. "org.lw/chapter4/lib2" // 使用.可以不用包名, 直接调用API, 注意不同包有相同的API接口会发生冲突
	
)

func main() {
	// lib1.Lib1Test()
	// lib2.Lib2Test()
	// mylib2.Lib2Test()
	Lib2Test()
}

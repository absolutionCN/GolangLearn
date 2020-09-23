package main

var (
	a int
	b string
	c []float32
	d func() bool
	e struct {
		x int
	}
)

// 名字 ：= 表达式
// （定义变量，同时显示初始化）
// 不能提供数据类型
// 只能在函数内部使用

// 简短变量声明语句也可以用来声明和初始化一组变量
//i, j := 0, 1

func main() {
	x := 100
	a, s := 1, "abc"
}

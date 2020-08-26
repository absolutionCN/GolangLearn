package main

import "fmt"

//func main() {
//	// 准备一个字符串
//	str := "hello world"
//
//	// 创建一个匿名函数
//	foo := func() {
//
//		// 匿名函数中访问str
//		str = "hello dude"
//	}
//	foo()
//}
//
//func Accumulate(value int) func() int {
//	// 返回一个闭包
//	return func() int {
//		//累加
//		value++
//		return value
//	}
//}
//
//func main() {
//	// 创建一个累加器，初始值为1
//	accumulate := Accumulate(1)
//	// 累加并打印
//	fmt.Println(accumulate())
//
//	fmt.Println(accumulate())
//
//	// 打印累加器的函数地址
//	fmt.Printf("%p\n", &accumulate)
//
//	// 创建一个累加器，初始值为1
//	accumulate2 := Accumulate(1)
//
//	//累加1并打印
//	fmt.Println(accumulate2())
//
//	//打印累加器的函数地址
//	fmt.Printf("%p\n", &accumulate2)
//
//}

//闭包实现生成器
func playerGen(name string) func() (string, int) {
	// 血量一直为150
	hp := 150

	// 返回创建的闭包
	return func() (string, int) {
		// 将变量引用到闭包中
		return name, hp
	}
}

func main() {
	// 创建一个玩家生成器
	generator := playerGen("hign noon")

	// 返回玩家的名字和血量
	name, hp := generator()

	//打印值
	fmt.Println(name, hp)
}

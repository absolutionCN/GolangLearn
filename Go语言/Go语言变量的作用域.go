package Go语言

import "fmt"

//func main() {
//
//	//声明局部变量a 和 b并赋值
//	var a int = 3
//	var b int = 4
//	//声明局部变量c并计算a和b的和
//	var c int
//	c = a + b
//	fmt.Println(c)
//
//}

//全局变量

//声明全局变量
//var c int
//
//func main() {
//	var a, b int
//	//初始化参数
//	a = 3
//	b = 4
//	c = a + b
//
//	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
//}

//Go语言程序中全局变量与局部变量名称可以相同，但是函数体内的局部变量会被优先考虑
// 声明全局变量
//var a float32 = 3.14
//
//func main() {
//	//声明局部变量
//	var a int = 3
//	fmt.Println(a)
//}

//全局变量a
var a int = 13

func main() {
	//局部变量a和b
	var a int = 3
	var b int = 4

	fmt.Printf("main() 函数中 a = %d\n", a)
	fmt.Printf("main() 函数中 b = %d\n", b)

	c := sum(a, b)

	fmt.Printf("main() 函数中 c = %d\n", c)
}

func sum(a, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)
	num := a + b
	return num
}

package Go语言

import "fmt"

//// 定义时调用匿名函数
//func(data int) {
//	fmt.Println("hello", data)
//}(100)
// 第3行}后的(100)，表示对匿名函数进行调用，传递参数为100

// 将匿名函数值赋给变量
//将匿名函数体保存到f()中
//f := func(data int) {
//	fmt.Println("hello", data)
//}
////使用f()调用
//f(100)

//匿名函数用作回调函数

////遍历切片的每个元素，通过给定函数进行元素访问
//func visit(list []int, f func(int)) {
//	for _, v := range list {
//		f(v)
//	}
//}
//func main() {
//	//使用匿名函数打印切片内容
//	visit([]int{1, 2, 3, 4}, func(v int) {
//		fmt.Println(v)
//	})
//}

// 匿名函数
//实例1
//func main() {
//	f := func(){
//		fmt.Println("hello world")
//	}
//	f() // hello world
//	fmt.Printf("%T\n", f) // 打印 func()
//}

//// 带参数
//func main() {
//	f := func(args string) {
//		fmt.Println(args)
//	}
//	f("hello world") // hello world
//	//或
//	(func(args string){
//		fmt.Println(args)
//	})("hello world")//hello world
//	//或
//	func(args string){
//		fmt.Println(args)
//	}("hello world")
//}
//// 带返回值
//func main() {
//	f := func()string{
//		return "hello world"
//	}
//	a := f()
//	fmt.Println(a)
//}

//// 多个匿名函数
//func main() {
//	f1, f2 := F(1, 2)
//	fmt.Println(f1(4))
//	fmt.Println(f2())
//}
//func F(x, y int)(func(int)int, func()int){
//	f1 := func(z int) int{
//		return (x + y) * z / 2
//	}
//	f2 := func() int {
//		return 2 * (x + y)
//	}
//	return f1, f2
//}
//闭包
//1
//func main() {
//	a := Fun()
//	b := a("hello ")
//	c := a("hello ")
//	fmt.Println(b)
//	fmt.Println(c)
//}
//func Fun() func(string) string{
//	a := "world"
//	return func(args string) string{
//		a += args
//		return a
//	}
//}
//2
//func main() {
//	a := Fun()
//	d := Fun()
//	b := a("hello")
//	c := a("hello")
//	e := d("hello")
//	f := d("hello")
//	fmt.Println(b)
//	fmt.Println(c)
//	fmt.Println(e)
//	fmt.Println(f)
//}
//func Fun() func(string) string{
//	a := "world"
//	return func(args string) string{
//		a += args
//		return a
//	}
//}
//3
//func main() {
//	a := F()
//	a[0]()
//	a[1]()
//	a[2]()
//}
//func F() []func(){
//	b := make([]func(), 3, 3)
//	for i := 0; i < 3; i ++ {
//		b[i] = func(){
//			fmt.Println(&i, i)
//		}
//	}
//	return b
//}
// 避免闭包参数不变
//func main() {
//	a := F()
//	a[0]()
//	a[1]()
//	a[2]()
//}
//func F() []func() {
//	b := make([]func(),3,3)
//	for i := 0; i < 3; i++{
//		b[i] = (func(j int) func(){
//			return func(){
//				fmt.Println(&j, j)
//			}
//		})(i)
//	}
//	return b
//}

//func main() {
//	fmt.Println(F())
//}
//
//func F() (r int){
//	defer func(){
//		r ++
//	}()
//	return 1
//}
//递归函数
//
//func F(i int) int{
//	if i <= 1{
//		return 1
//	}
//	return i * F(i-1)
//}
//func main(){
//	var i int = 3
//	fmt.Println(i , F(i))
//}
//斐波那契数列
func fibonaci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
}

func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\n", fibonaci(i))
	}
}

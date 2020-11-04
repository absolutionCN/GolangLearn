package Go语言

import (
	"fmt"
	"time"
)

func test1() {
	start := time.Now() // 获取当前时间

	sum := 0
	for i := 0; i < 100000000; i++ {
		sum++
	}
	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}

//func main() {
//	test()
//}

// 使用time.Now().Sub()获取函数运行时间

func test2() {
	start := time.Now()
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum++
	}
	elapsed := time.Now().Sub(start)
	fmt.Println("该函数执行完成耗时:", elapsed)
}

func main() {
	test1()
	test2()
}

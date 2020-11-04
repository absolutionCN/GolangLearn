package Go语言

import "fmt"

// 使用键值对初始化结构体
//type People struct {
//	name  string
//	child *People
//}
//
//func main() {
//	relation  := &People{
//		name: "爷爷",
//		child: &People{
//			name: "爸爸",
//			child: &People{
//				name: "我",
//			},
//		},
//	}
//	fmt.Println(relation)
//}

//使用多个值的列表初始化结构体
//
//type Address struct {
//	Province    string
//	City        string
//	ZipCode     int
//	PhoneNumber string
//}
//
//func main() {
//	addr := Address{
//		"四川",
//		"成都",
//		610000,
//		"0",
//	}
//	fmt.Println(addr)
//}

//初始化匿名结构体

//打印消息类型，传入匿名结构体
func printMsgType(msg *struct {
	id   int
	data string
}) {
	// 使用动词%T打印msg的类型
	fmt.Printf("%T\n", msg)
}

func main() {
	// 实例化一个匿名结构体
	msg := &struct { // 定义部分
		id   int
		data string
	}{ // 值初始化部分
		1024,
		"hello",
	}
	printMsgType(msg)
}

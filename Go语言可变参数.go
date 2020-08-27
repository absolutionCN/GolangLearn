package main

import "fmt"

//明日香
//func main() {
//	myfunc(1,2,3,4,5,6,7,8,90)
//}
//func myfunc(args ...int) {
//	for _, arg := range args{
//		fmt.Println(arg)
//	}
//}

//func アスカ(args ...interface{}) {
//	for _, arg := range args {
//		switch arg.(type) {
//		case int:
//			fmt.Println(arg, "is an int value")
//		case string:
//			fmt.Println(arg, "is a string value")
//		case int64:
//			fmt.Println(arg, "is an int64 value")
//		default:
//			fmt.Println(arg, "is an unknown type")
//		}
//	}
//}
//
//func main() {
//	var v1 int = 1
//	var v2 int64 = 234
//	var v3 string = "hello"
//	var v4 float32 = 1.234
//	アスカ(v1, v2, v3, v4)
//}

////定义一个函数，参数数量为0~n，类型约束为字符串
//func アスカ(slist ...string) string{
//	// 定义一个字节缓冲，快速地连接字符串
//	var b bytes.Buffer
//	// 遍历可变参数列表slist, 类型为[]string
//	for _, s := range slist{
//		// 将遍历出的字符串连续写入字节数组
//		b.WriteString(s)
//	}
//	// 将连接好的字节数组转换为字符串输并输出
//	return b.String()
//}
//
//func main() {
//	//输入3个字符串将他们连成一个字符串  ayanamirei　　いかりしんじ
//	fmt.Println(アスカ("アスカ", "綾波レイ", "碇シンジ"))
//	fmt.Println(アスカ("02", "00", "01"))
//}

//获得可变参数类型——获得每一个参数的类型
//func printTypeValue(slist ...interface{}) string {
//	// 字节缓冲作为快速字符串连接
//	var b bytes.Buffer
//
//	// 遍历参数
//	for _, s := range slist {
//		// 将interface{}类型格式化为字符串
//		str := fmt.Sprintf("%v", s)
//
//		// 类型的字符串描述
//		var typeString string
//
//		// 对s进行类型断言
//		switch s.(type) {
//		case bool: // 当s为布尔类型时
//			typeString = "bool"
//		case string: //当s为字符串类型时
//			typeString = "string"
//		case int: // 当s为整型类型时
//			typeString = "int"
//		}
//		// 写字符串前缀
//		b.WriteString("value:")
//		// 写入值
//		b.WriteString(str)
//		//写入类型前缀
//		b.WriteString(typeString)
//		//写入换行符
//		b.WriteString("\n")
//	}
//	return b.String()
//}
//
//func main() {
//	// 将不同类型的变量通过printTypeValue()打印出来
//	fmt.Println(printTypeValue(100, "str", true))
//}

//在多个可变参数函数中传递参数
// 实际打印的函数
func rawPrint(rawList ...interface{}) {

	//遍历可变参数切片
	for _, a := range rawList {
		//打印参数
		fmt.Println(a)
	}
}

//打印函数封装
func print(slist ...interface{}) {
	//将slist可变参数切片完整传递给下一个函数
	rawPrint(slist...)
	//rawPrint("fmt", slist)
}

func main() {
	print(1, 2, 3)
}

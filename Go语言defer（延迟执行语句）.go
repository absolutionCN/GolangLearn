package main

import "os"

//func main() {
//	fmt.Println("defer begin")
//
//	// 将defer放入延迟调用栈
//	defer fmt.Println(1)
//
//	defer fmt.Println(2)
//
//	// 最后一个放入,位于栈顶，最先调用
//	defer fmt.Println(3)
//
//	fmt.Println("defer end")
//
//}

////使用延迟并发解锁
//var (
//	//
//	valueByKey = make(map[string]int)
//
//	//保证使用映射时的并发安全的互斥锁
//	valueByKeyGuard sync.Mutex
//)
//
////根据键读取值
//func readValue(key string) int{
//	// 对共享资源加锁
//	valueByKeyGuard.Lock()
//	//取值
//	v := valueByKey[key]
//	//对共享资源解锁
//	valueByKeyGuard.Unlock()
//	//返回值
//	return v
//}

//// 使用defer语句对上面语句进行简化
//func readValue(key string) int {
//	valueByKeyGuard.Lock()
//
//	//def后面的语句不会马上调用，而是延迟到函数结束时调用
//	defer valueByKeyGuard.Unlock()
//
//	return valueByKey[key]
//}

//使用延迟释放文件句柄
////根据文件名查询其大小
//func fileSize(filename string) int64 {
//
//	// 根据文件名打开文件，返回文件句柄和错误
//	f, err := os.Open(filename)
//
//	// 如果打开时发生错误，返回文件大小为0
//	if err != nil{
//		return 0
//	}
//
//	//取文件状态信息
//	info, err:= f.Stat()
//
//	// 如果取信息时发生错误，关闭文件并返回文件大小为0
//	if err != nil {
//		f.Close()
//		return 0
//	}
//
//	//取文件大小
//	size := info.Size()
//
//	// 关闭文件
//	f.Close()
//
//	// 返回文件大小
//	return size
//}

// 对上面代码进行简化
func fileSize(filename string) int64 {
	f, err := os.Open(filename)

	if err != nil {
		return 0
	}
	// 延迟调用Close，此时Close不会被调用
	defer f.Close()

	info, err := f.Stat()

	if err != nil {
		return 0
	}
	size := info.Size()

	// defer机制触发，调用Close关闭文件
	return size

}

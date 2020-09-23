package main

// 声明一个结构体
type class struct {
}

//// 给结构体添加Do方法
//func (c *class) Do(v int){
//
//	fmt.Println("call method do:", v)
//
//}
//
//// 普通函数的Do
//func funcDo(v int){
//
//	fmt.Println("call function do:", v)
//
//}
//
//func main() {
//
//	// 声明一个函数回调
//	var delegate func(int)
//
//	// 创建结构体实例
//	c := new(class)
//
//	// 将回调设为c的Do方法
//	delegate = c.Do
//
//	// 调用
//	delegate(100)
//
//	// 将回调设为普通函数
//	delegate = funcDo
//
//	// 调用
//	delegate(100)
//
//}

// 实例化一个通过字符串映射函数切片的map
var eventByName = make(map[string][]func(interface{}))

// 注册事件，提供事件名和回调函数
func RegisterEvent(name string, callback func(interface{})) {

	// 通过名字查找事件列表
	list := eventByName[name]

	// 在列表切片中添加函数
	list = append(list, callback)

	// 将修改的事件列表切片保存回去
	eventByName[name] = list

}

// 调用事件
func CallEvent(name string, param interface{}) {
	// 通过名字找到事件列表
	list := eventByName[name]

	// 遍历这个事件的所有回调
	for _, callback := range list {
		//传入参数调用回调
		callback(param)
	}
}

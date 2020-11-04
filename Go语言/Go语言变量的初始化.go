package Go语言

// 变量初始化的标准格式
// var 变量名 类型 = 表达式
////var hp int = 100
//var hp = 100
//
//// 编译器根据右值推导变量类型完成初始化的样子
//var attack = 40
//var defence = 20
//var damageRate float32 = 0.17
//var damage = float32(attack - defence) * damageRate
//
//func main() {
//	fmt.Println(damage)
//}


// 有新的变量在左值中出现，其他变量名可以重复声明
conn, err := net.Dial("tcp", "127.0.0.1")
conn2, err := net.Dial("tcp", "127.0.0.1")

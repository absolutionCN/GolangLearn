package main

//type Point struct {
//	X int
//	Y int
//}
//
//type Color struct {
//	R, G, B byte
//}
//func main() {
//	type Point struct {
//		X int
//		Y int
//	}
//
//	// 创建指针类型的结构体
//	type Player struct {
//		Name string
//		HealthPoint int
//		MagicPoint int
//	}
//
//	tank := new(Player)
//	tank.Name = "Canon"
//	tank.HealthPoint = 300
//
//	//取结构体的地址实例化
//	//ins := &T{}
//	type Command struct {
//		Name string // 指令名称
//		Var *int // 指令绑定的变量
//		Comment string // 指令的注释
//	}
//
//	var version int = 1
//	cmd := &Command{}
//	cmd.Name = "version"
//	cmd.Var = &version
//	cmd.Comment = "show version"
//
//
//	var p Point
//	p.X = 10
//	p.Y = 20
//}

// 取地址实例化是最广泛的一种结构体实例化方式，可以使用函数封装上面的初始化过程
func newCommand(name string, varref *int, comment string) *Command {
	return &Command{
		Name:    name,
		Var:     varref,
		Comment: comment,
	}
}
func main() {
	cmd = newCommand(
		"version",
		&version,
		"show version",
	)

}

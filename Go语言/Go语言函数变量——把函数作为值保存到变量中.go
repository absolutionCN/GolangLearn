package Go语言

import "fmt"

func fire() {
	fmt.Println("fire")
}

func main() {
	var f func()
	f = fire
	f()
}

package main

import "fmt"

var a = 100
var b = 200

//func main() {
//	a = a ^ b
//	b = b ^ a
//	a = a ^ b
//
//	fmt.Println(a, b)
//}
func main() {
	b, a = a, b
	fmt.Println(a, b)
}

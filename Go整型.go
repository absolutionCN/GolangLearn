package main

import "fmt"

var u uint8 = 255
var i int8 = 127

func main() {
	fmt.Println(u, u+1, u*u)

	fmt.Println(i, i+1, i*i) // "127 -128 1"
}

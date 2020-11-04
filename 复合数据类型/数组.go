package main

import "fmt"

//var a [3]int
//
//func main() {
//	fmt.Println(a[0])
//	fmt.Println(a[len(a)-1])
//
//	for i, v := range a {
//		fmt.Printf("%d %d\n", i ,v )
//	}
//}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	//data := []string{"one", "", "three"}
	//fmt.Println(nonempty2(data))
	fmt.Println(nonempty2([]string{"one", "", "three", "four"}))

}

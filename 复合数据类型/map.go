package main

import (
	"bufio"
	"fmt"
	"os"
)

//var names []string

//var ages map[string]int

//func main() {
//	//names := make([]string, 0 , len(ages))
//	//ages := map[string]int{
//	//	"alice":   31,
//	//	"charlie": 34,
//	//}
//	//for name := range ages {
//	//	names = append(names, name)
//	//}
//	//sort.Strings(names)
//	//for _, name:= range names {
//	//	fmt.Printf("%s\t%d\n", name, ages[name])
//	//}
//	//ages["carol"] = 21
//	//if age, ok := ages["bob"]; !ok { /* ... */ }
//}

//func equal(x, y map[string]int) bool {
//	if len(x) != len(y) {
//		return false
//	}
//	for k, xv := range x {
//		if yv, ok := y[k]; !ok || yv != xv {
//			return false
//		}
//	}
//	return true
//}
//
//func main() {
//	a := equal(map[string]int{"A": 0}, map[string]int{"B": 42})
//
//	fmt.Println(a)
//}

func main() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup:%v\n", err)
		os.Exit(1)
	}
}

package main

import "fmt"

//type Employee struct {
//
//		ID           int
//		Name         string
//		Address      string
//		Dob          time.Time
//		Position     string
//		Salary       int
//		ManagerID    int
//}
//
//var diblert Employee
//
//
//func main() {
//	diblert.Salary -= 5000
//
//	position := &diblert.Position
//	*position = "Senior" + *position
//	var employeeOfTheMonth *Employee = &dilbert
//	employeeOfTheMonth.Position += " (proactive team player)"
//}

//type tree struct {
//	value       int
//	left, right *tree
//}
//
//// Sort sorts values in place.
//func Sort(values []int) {
//	var root *tree
//	for _, v := range values {
//		root = add(root, v)
//	}
//	appendValues(values[:0], root)
//}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
//func appendValues(values []int, t *tree) []int {
//	if t != nil {
//		values = appendValues(values, t.left)
//		values = append(values, t.value)
//		values = appendValues(values, t.right)
//	}
//	return values
//}
//
//func add(t *tree, value int) *tree {
//	if t == nil {
//		// Equivalent to return &tree{value: value}.
//		t = new(tree)
//		t.value = value
//		return t
//	}
//	if value < t.value {
//		t.left = add(t.left, value)
//	} else {
//		t.right = add(t.right, value)
//	}
//	return t
//}

//type Ponit struct{x, y int}
//p := Point{1, 2}

type Point struct {
	x, y int
}

type address struct {
	hostname string
	port     int
}

func main() {
	p := Point{1, 2}
	q := Point{2, 1}
	fmt.Println(p.x == q.x && p.y == q.y)
	fmt.Println(p == q)

	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
}

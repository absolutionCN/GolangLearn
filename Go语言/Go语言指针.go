package Go语言

import (
	"fmt"
)

//var n = flag.Bool("n", false, "omit trailing newline")
//var sep = flag.String("s", " ", "separator")
//
//func main() {
//	flag.Parse()
//	fmt.Print(strings.Join(flag.Args(), *sep))
//	if !*n{
//		fmt.Println()
//	}
//}

//func main() {
//	p := new(int)
//	fmt.Println(*p)
//	*p = 2
//	fmt.Println(*p)
//}

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC))
	//fmt.Printf("&g\n", boilingF-FreezingC)
}

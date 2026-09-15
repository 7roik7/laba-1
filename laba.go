package main

import (
	"fmt"
	"time"
)

func sumAndDiff(a, b float64) (float64, float64) { return a + b, a - b }
func average(a, b, c float64) float64            { return (a + b + c) / 3 }

func main() {
	// 1
	fmt.Println("Дата:", time.Now().Format("02.01.2006 15:04:05"))

	// 2
	var i int = 42
	var f float64 = 3.14
	var s string = "Go"
	var b bool = true
	fmt.Println(i, f, s, b)

	// 3
	x, y, z, w := 1, 2.5, "text", false
	fmt.Println(x, y, z, w)

	// 4
	a, c := 17, 5
	fmt.Println(a+c, a-c, a*c, a/c, a%c)

	// 5
	s1, d1 := sumAndDiff(7.5, 2.3)
	fmt.Println(s1, d1)

	// 6
	fmt.Println(average(4, 8, 15))
}

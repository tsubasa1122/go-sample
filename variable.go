package main

import "fmt"

func main() {
	var i int = 1
	var f64 float64 = 1.2
	var s string = "test"
	var t, f bool = true, false
	fmt.Println(i, f64, s, t, f)

	xi := 1
	xf64 := 1.2
	fmt.Println(xi)
	fmt.Printf("%T\n", xf64)
}

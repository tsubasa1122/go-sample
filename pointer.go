package main

import "fmt"

func inc(p *int) {
	*p++
}

func main() {
	i := 1
	inc(&i)
	inc(&i)
	inc(&i)
	fmt.Println(i)

	a := [3]string{"Apple", "Banana", "Cherry"}
	p := &a
	fmt.Println(a[1])
	fmt.Println(p[1])

	i2 := 5
	ip := &i2
	fmt.Println("type=%T, address=%p\n", ip, ip)
}

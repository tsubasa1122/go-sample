package main

import "fmt"

func main() {
	s := make([]int, 10, 20)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a[:])

	b := []int{1, 2, 3}
	b = append(b, 4)
	fmt.Println(b)

	c := make([]int, 0, 0)
	fmt.Printf("len=%d, cap=%d\n", len(c), cap(c))

	c = append(c, 1)
	fmt.Printf("len=%d, cap=%d\n", len(c), cap(c))

	c = append(c, []int{2, 3, 4}...)
	fmt.Printf("len=%d, cap%d\n", len(c), cap(c))

	c = append(c, 5)
	fmt.Printf("len=%d, cap=%d\n", len(c), cap(c))

	c = append(s, 6, 7, 8, 9)
	fmt.Printf("len=%d, cap=%d\n", len(c), cap(c))

	d := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := d[2:4]
	fmt.Println(s1)
	fmt.Printf("len=%d, cap=%d\n", len(s1), cap(s1))

	s2 := d[2:4:4]
	fmt.Println(s2)
	fmt.Printf("len=%d, cap=%d\n", len(s2), cap(s2))

	s3 := d[2:4:6]
	fmt.Println(s3)
	fmt.Printf("len=%d, cap=%d\n", len(s3), cap(s3))

	s4 := []string{"Apple", "Banana", "Cherry"}
	for i, v := range s4 {
		fmt.Printf("[%d] => %s\n", i, v)
	}
}

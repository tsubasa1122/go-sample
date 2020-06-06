package main

import "fmt"

func main() {
	m := make(map[int]string)
	m[1] = "US"
	m[81] = "Japan"
	m[86] = "China"

	fmt.Println(m)

	m2 := map[int][]int{
		1: []int{1},
		2: []int{1, 2},
		3: []int{1, 2, 3},
	}

	fmt.Println(m2)

	m3 := map[int]string{1: "A", 2: "B", 3: "C"}

	s, ok := m3[1]
	fmt.Println(s)
	fmt.Println(ok)

	s1, ok := m3[9]
	fmt.Println(s1)
	fmt.Println(ok)

	_, ok2 := m3[3]
	fmt.Println(ok2)

	m4 := map[int][]int{
		1: []int{1},
		2: []int{1, 2},
		3: []int{1, 2, 3},
	}

	fmt.Println(m4)

	m5 := map[int]map[float64]string{
		1: {3.14: "円周率"},
	}

	fmt.Println(m5)
}

package main

import "fmt"

func thirdPartyConnectionDB() {
	panic("Unable to connect database!")
}

func save() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectionDB()
}

func main() {
	save()
	fmt.Println("ok?")
}

package main

type Person struct {
	name string
	age  int
}

func (p *Person) SetPerson(name string, age int) {
	p.name = name
	p.age = age
}

func (p *Person) GetPerson() (string, int) {
	p.name
}

func main() {
	var p1 Person
	p1.SetPerson("Yamada", 26)
	name, age = p1.GetPerson()
}

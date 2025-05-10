package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	var a int = 5

	var b int = 10
	fmt.Println("Sum:", a+b)
	fmt.Println("Difference:", a-b)
	a1 := 10
	fmt.Println("a1:", a1)
	var typ bool = true
	fmt.Println("val:", typ)

	var t1 []int = []int{1, 2, 3}
	fmt.Print("t1:", t1)
	fmt.Println("t1[0]:", t1[0])

	fmt.Println("t1 after append:", t1)

	var per1 Person = Person{id: 1, age: 20, name: "John"}
	fmt.Println("per1:", per1)

	var stud1 Student = Student{name: "Alice", rollno: 101}
	fmt.Println("stud1:", stud1)
	fmt.Println("stud1.name:", stud1.name)

	stud2 := &stud1
	stud2.setName("Bob")
	fmt.Printf("stud1 after setName: %v", stud2.name)
}

type Person struct {
	id   int
	name string
	age  int
}

type Student struct {
	name   string
	rollno int
}

func (P Person) getName() string {
	return "age" + P.name
}

func (P Person) getId() int {
	return P.id
}

func (S *Student) setName(name string) {
	S.name = name
}

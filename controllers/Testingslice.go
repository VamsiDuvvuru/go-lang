package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var student1 student = student{
		Name:    "John Doe",
		Age:     20,
		Grade:   "A",
		Address: "123 Main St, Springfield",
		Phone:   "123-456-7890",
	}

	fmt.Println("Student1 Details:", student1)
	student2 := &student1
	student2.updateStudentName("Jane larcy")
	fmt.Println("Student2 Details:", *student2)

	var students []student = []student{
		{
			Name:    "Alice Smith",
			Age:     22,
			Grade:   "B",
			Address: "456 Elm St, Springfield",
			Phone:   "987-654-3210",
		},
		{
			Name:    "Bob Johnson",
			Age:     21,
			Grade:   "C",
			Address: "789 Oak St, Springfield",
			Phone:   "555-123-4567",
		},
		{
			Name:    "Charlie Brown",
			Age:     23,
			Grade:   "A",
			Address: "321 Pine St, Springfield",
			Phone:   "444-987-6543",
		},
	}

	name := "Bob Johnson"
	for i, stud := range students {
		if stud.Name == name {
			students = append(students[:i], students[i+1:]...)
		}
	}
	fmt.Println("Updated Students List:", students)

	for i, s := range students {
		if s.Name == "Alice Smith" {
			students[i].updateStudentName("Alice Johnson*")
			break
		}
	}

	fmt.Println("****Updated Students List:*****", students)
	data, err1 := json.MarshalIndent(students, "", "  ")
	if err1 != nil {
		fmt.Println("Error marshaling to JSON:", err1)
		return
	}

	err := os.WriteFile("students.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	data, err = os.ReadFile("students.json")
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
	var readStudents []student
	err = json.Unmarshal(data, &readStudents)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	fmt.Println("Read Students from JSON file:", readStudents)
}

type student struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Grade   string `json:"grade"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func (s *student) updateStudentName(name string) {
	s.Name = name
}

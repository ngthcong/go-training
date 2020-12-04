package main

import (
	"bufio"
	"fmt"
	"os"
	"school-mongo/repositories"
)

var (
	scanner *bufio.Scanner
)

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	fmt.Println("-----School Management-----")
	fmt.Println("1 : Class Management")
	fmt.Println("2 : Discipline Management")
	fmt.Println("3 : Teacher Management")
	fmt.Println("4 : Student Management")
	fmt.Println("5 : Exit")
	fmt.Println("---------------------------")
	fmt.Println("Selection: ")

	scanner.Scan()
	text := scanner.Text()

	fmt.Println(text)
	switch text {
	case "1":
		StudentManager()
	case "2":
		StudentManager()
	case "3":
		StudentManager()
	case "4":
		StudentManager()
	case "5":
		StudentManager()
	case "6":
		os.Exit(0)
	default:
		fmt.Println("Please select from 1 to 6")
	}

}

func StudentManager() {
	fmt.Println("---------------------------")
	fmt.Println("1 : Add student")
	fmt.Println("2 : Edit student")
	fmt.Println("3 : Delete student")
	fmt.Println("4 : Return")
	fmt.Println("---------------------------")
	fmt.Println("Selection: ")
	scanner.Scan()
	text := scanner.Text()
	switch text {
	case "1":
		AddStudent()
	case "2":
		EditStudent()
	case "3":
		DeleteTeacher()
	case "4":
		main()
	default:
		fmt.Println("Please select from 1 to 4")
	}
}
func AddStudent() {
	fmt.Println("---------------------------")
	fmt.Println("Student id: ")
	scanner.Scan()
	id := scanner.Text()
	fmt.Println("Student name: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Println("Class id: ")
	scanner.Scan()
	class := scanner.Text()
	newStudent := repositories.Student{
		Id:      id,
		Name:    name,
		ClassID: class,
	}

	fmt.Println("---------------------------")
	fmt.Printf("Student name: %s Class id: %s", newStudent.Name, newStudent.ClassID)
}
func EditStudent() {
	fmt.Println("---------------------------")
	fmt.Println("Student id: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Class id: ")
	scanner.Scan()
	class := scanner.Text()

	fmt.Println("---------------------------")
	fmt.Printf("Student name: %s Class id: %s", name, class)
}
func DeleteStudent() {
	fmt.Println("---------------------------")
	fmt.Println("Student id: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Class id: ")
	scanner.Scan()
	class := scanner.Text()

	fmt.Println("---------------------------")
	fmt.Printf("Student name: %s Class id: %s", name, class)
}

func TeacherManager() {
	fmt.Println("---------------------------")
	fmt.Println("1 : Add teacher")
	fmt.Println("2 : Edit teacher")
	fmt.Println("3 : Delete teacher")
	fmt.Println("4 : Return")
	fmt.Println("---------------------------")
	fmt.Println("Selection: ")
	scanner.Scan()
	text := scanner.Text()
	switch text {
	case "1":
		AddTeacher()
	case "2":
		EditTeacher()
	case "3":
		DeleteTeacher()
	case "4":
		main()
	default:
		fmt.Println("Please select from 1 to 4")
	}
}
func AddTeacher() {
	fmt.Println("---------------------------")
	fmt.Println("Teacher name: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Println("Class id: ")
	scanner.Scan()
	class := scanner.Text()

	fmt.Println("---------------------------")
	fmt.Printf("Teacher name: %s Class id: %s", name, class)
}
func EditTeacher() {
	fmt.Println("---------------------------")
	fmt.Println("Teacher id: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Class id: ")
	scanner.Scan()
	class := scanner.Text()

	fmt.Println("---------------------------")
	fmt.Printf("Teacher name: %s Class id: %s", name, class)
}
func DeleteTeacher() {
	fmt.Println("---------------------------")
	fmt.Println("Teacher id: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Class id: ")
	scanner.Scan()
	class := scanner.Text()

	fmt.Println("---------------------------")
	fmt.Printf("Teacher name: %s Class id: %s", name, class)
}

package teacher

import (
	"bufio"
	"fmt"
	"school-mongo/internal/model"

	"school-mongo/internal/services/teacher"
)

var (
	scanner *bufio.Scanner
)

type (
	handler struct {
		s teacher.Service
	}
)

func New(service teacher.Service) *handler {
	return &handler{s: service}
}

func (h *handler) Manage() bool {
	fmt.Println("---------------------------")
	fmt.Println("1: Add teacher")
	fmt.Println("2: Edit teacher")
	fmt.Println("3: Delete teacher")
	fmt.Println("4: Return")
	fmt.Println("---------------------------")
	fmt.Println("Selection: ")
	scanner.Scan()
	text := scanner.Text()
	switch text {
	case "1":
		h.Add()
	case "2":
		h.Edit()
	case "3":
		h.Delete()
	case "4":
		return true
	default:
		fmt.Println("Please select from 1 to 4")
	}
	h.Manage()
	return false
}
func (h *handler) Add() {
	fmt.Println("---------------------------")
	fmt.Println("Teacher ID: ")
	scanner.Scan()
	id := scanner.Text()
	fmt.Println("Teacher name: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Println("Class id: ")
	scanner.Scan()
	class := scanner.Text()
	fmt.Println("Disciplines(separate by a space):  ")
	scanner.Scan()
	disciplines := scanner.Text()


	fmt.Println("---------------------------")



	teacher := model.Teacher{
		Id:           id,
		Name:         name,
		ClassID:      class,
		Disciplines: "",
	}
	fmt.Println(teacher.ToString())
	h.s.Add(teacher)
	h.Manage()
}
func (h *handler) Delete() {

	h.Manage()
}
func (h *handler) Edit() {
	fmt.Println("---------------------------")

	fmt.Println("Teacher id: ")
	scanner.Scan()
	id := scanner.Text()
	teacher, err := h.s.Get(id)
	if err != nil {
		fmt.Printf("Teacher with ID %s not found \n", id)
	}
	fmt.Println("Teacher name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Class ID: ")
	scanner.Scan()
	class := scanner.Text()

	teacher := model.Teacher{
		Id:          "",
		Name:        "",
		ClassID:     "",
		Disciplines: "",
	}
	h.s.Update()
	fmt.Println("---------------------------")
	fmt.Printf("Teacher name: %s Class id: %s", name, class)
	h.Manage()
}

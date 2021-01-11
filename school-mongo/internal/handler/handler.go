package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"school-mongo/internal/model"
	"school-mongo/internal/services"
)

type handler struct {
	sc *bufio.Scanner
	s  services.Service
}

// Handler ...
type Handler interface {
	InitProgram() error
	ClassManager()
	DisciplineManager()
	TeacherManager()
	StudentManager()
}

//New ...
func New(scanner *bufio.Scanner, service services.Service) handler {
	return handler{
		sc: scanner,
		s:  service,
	}
}

func (h *handler) InitProgram() {
	fmt.Println("-----School Management-----")
	fmt.Println("1: Class Management")
	fmt.Println("2: Discipline Management")
	fmt.Println("3: Teacher Management")
	fmt.Println("4: Student Management")
	fmt.Println("5: Exit")
	fmt.Println("---------------------------")
	fmt.Println("Selection: ")
	h.sc.Scan()
	text := h.sc.Text()

	fmt.Println(text)
	switch text {
	case "1":
		h.ClassManager()
	case "2":
		h.DisciplineManager()
	case "3":
		h.TeacherManager()
	case "4":
		h.StudentManager()
	case "5":
		os.Exit(0)
	default:
		fmt.Println("Please select from 1 to 6")
	}
	h.InitProgram()
}

func (h *handler) ClassManager() {
	fmt.Println("---------------------------")
	fmt.Println("1: Add class")
	fmt.Println("2: Edit class")
	fmt.Println("3: Delete class")
	fmt.Println("4: Return")
	fmt.Println("---------------------------")
	fmt.Println("Selection: ")

	h.sc.Scan()
	text := h.sc.Text()
	switch text {
	case "1":
		h.AddClass()
	case "2":
		h.EditClass()
	case "3":
		h.DeleteClass()
	case "4":
		h.InitProgram()
	default:
		fmt.Println("Please select from 1 to 4")
	}
	h.ClassManager()
}

func (h *handler) AddClass() {
	fmt.Println("---------------------------")
	fmt.Println("Class ID: ")
	h.sc.Scan()
	id := h.sc.Text()
	fmt.Println("Name: ")
	h.sc.Scan()
	name := h.sc.Text()
	fmt.Println("Teachers (separate by a \",\"): ")
	h.sc.Scan()
	teachers := h.sc.Text()

	fmt.Println("---------------------------")
	strings.Split(teachers, ",")
	class := model.Class{
		ID:       id,
		Name:     name,
		Teachers: teachers,
	}
	fmt.Println(class.ToString())
	if err := h.s.ClassService.Add(class); err != nil {
		panic(err)
	}
	fmt.Println("Add class success")
	h.ClassManager()
}

func (h *handler) DeleteClass() {
	fmt.Println("---------------------------")
	fmt.Println("Class id: ")
	h.sc.Scan()
	id := h.sc.Text()

	if err := h.s.ClassService.Delete(id); err != nil {
		panic(err)
	}

	fmt.Printf("Delete class with id: %s \n", id)
	h.ClassManager()
}

func (h *handler) EditClass() {
	fmt.Println("---------------------------")
	fmt.Println("Class id: ")
	h.sc.Scan()
	id := h.sc.Text()
	class, err := h.s.ClassService.Get(id)
	if err != nil {
		fmt.Printf("Class with ID %s not found \n", id)
		h.EditClass()
	}
	fmt.Println("Class name: ")
	h.sc.Scan()
	name := h.sc.Text()

	fmt.Println("Teachers (separate by a \",\"): ")
	h.sc.Scan()
	teachers := h.sc.Text()

	class.Name = name
	class.Teachers = teachers
	if err := h.s.ClassService.Update(class); err != nil {
		panic(err)
	}

	fmt.Println("---------------------------")
	fmt.Printf("Class name: %s, teachers: %s", name, teachers)
	h.ClassManager()
}

func (h *handler) DisciplineManager() {
	fmt.Println("---------------------------")
	fmt.Println("1: Add discipline")
	fmt.Println("2: Edit discipline")
	fmt.Println("3: Delete discipline")
	fmt.Println("4: Return")
	fmt.Println("---------------------------")
	fmt.Println("Selection: ")

	h.sc.Scan()
	text := h.sc.Text()
	switch text {
	case "1":
		h.AddClass()
	case "2":
		h.EditClass()
	case "3":
		h.DeleteClass()
	case "4":
		h.InitProgram()
	default:
		fmt.Println("Please select from 1 to 4")
	}
	h.DisciplineManager()
}

func (h *handler) AddDiscipline() {
	fmt.Println("---------------------------")
	fmt.Println("Discipline ID: ")
	h.sc.Scan()
	id := h.sc.Text()
	fmt.Println("Name: ")
	h.sc.Scan()
	name := h.sc.Text()
	fmt.Println("Lectures: ")
	h.sc.Scan()
	lectures := h.sc.Text()
	fmt.Println("Exercises: ")
	h.sc.Scan()
	exercises := h.sc.Text()

	fmt.Println("---------------------------")
	dis := model.Discipline{
		ID:        id,
		Name:      name,
		Lectures:  lectures,
		Exercises: exercises,
	}

	fmt.Println(dis.ToString())
	if err := h.s.DisciplineService.Add(dis); err != nil {
		panic(err)
	}
	fmt.Println("Add discipline success")
	h.DisciplineManager()
}

func (h *handler) DeleteDiscipline() {
	fmt.Println("---------------------------")
	fmt.Println("Discipline id: ")
	h.sc.Scan()
	id := h.sc.Text()

	if err := h.s.DisciplineService.Delete(id); err != nil {
		panic(err)
	}

	fmt.Printf("Delete discipline with id: %s \n", id)
	h.DisciplineManager()
}

func (h *handler) EditDiscipline() {
	fmt.Println("---------------------------")
	fmt.Println("Discipline id: ")
	h.sc.Scan()
	id := h.sc.Text()
	dis, err := h.s.DisciplineService.Get(id)
	if err != nil {
		fmt.Printf("Discipline with ID %s not found \n", id)
		h.EditDiscipline()
	}
	fmt.Println("Discipline name: ")
	h.sc.Scan()
	name := h.sc.Text()

	fmt.Println("Lectures: ")
	h.sc.Scan()
	lectures := h.sc.Text()
	fmt.Println("Exercises: ")
	h.sc.Scan()
	exercises := h.sc.Text()

	dis.Name = name
	dis.Lectures = lectures
	dis.Exercises = exercises
	if err := h.s.DisciplineService.Update(dis); err != nil {
		panic(err)
	}

	fmt.Println("---------------------------")
	fmt.Printf("Discipline name: %s, letures: %s, exercises: %s", name, lectures, exercises)
	h.DisciplineManager()
}

func (h *handler) TeacherManager() {
	fmt.Println("---------------------------")
	fmt.Println("1: Add teacher")
	fmt.Println("2: Edit teacher")
	fmt.Println("3: Delete teacher")
	fmt.Println("4: Return")
	fmt.Println("---------------------------")
	fmt.Println("Selection: ")

	h.sc.Scan()
	text := h.sc.Text()
	switch text {
	case "1":
		h.AddTeacher()
	case "2":
		h.EditTeacher()
	case "3":
		h.DeleteTeacher()
	case "4":
		h.InitProgram()
	default:
		fmt.Println("Please select from 1 to 4")
	}
	h.TeacherManager()
}

func (h *handler) AddTeacher() {
	fmt.Println("---------------------------")
	fmt.Println("Teacher ID: ")
	h.sc.Scan()
	id := h.sc.Text()
	fmt.Println("Name: ")
	h.sc.Scan()
	name := h.sc.Text()
	fmt.Println("Disciplines (separate by a \",\"): ")
	h.sc.Scan()
	disc := h.sc.Text()

	fmt.Println("---------------------------")
	teacher := model.Teacher{
		ID:          id,
		Name:        name,
		Disciplines: disc,
	}

	fmt.Println(teacher.ToString())
	if err := h.s.TeacherService.Add(teacher); err != nil {
		panic(err)
	}
	fmt.Println("Add teacher success")
	h.ClassManager()
}

func (h *handler) DeleteTeacher() {
	fmt.Println("---------------------------")
	fmt.Println("Teacher id: ")
	h.sc.Scan()
	id := h.sc.Text()

	if err := h.s.TeacherService.Delete(id); err != nil {
		panic(err)
	}

	fmt.Printf("Delete teacher with id: %s \n", id)
	h.ClassManager()
}

func (h *handler) EditTeacher() {
	fmt.Println("---------------------------")
	fmt.Println("Teacher id: ")
	h.sc.Scan()
	id := h.sc.Text()
	teacher, err := h.s.TeacherService.Get(id)
	if err != nil {
		fmt.Printf("Teacher with ID %s not found \n", id)
		h.EditClass()
	}
	fmt.Println("Teacher name: ")
	h.sc.Scan()
	name := h.sc.Text()

	fmt.Println("Disciplines: ")
	h.sc.Scan()
	disc := h.sc.Text()

	teacher.Name = name
	teacher.Disciplines = disc
	if err := h.s.TeacherService.Update(teacher); err != nil {
		panic(err)
	}

	fmt.Println("---------------------------")
	fmt.Printf("Teacher name: %s, discipline: %s", name, disc)
	h.TeacherManager()
}
func (h *handler) StudentManager() {
	fmt.Println("---------------------------")
	fmt.Println("1: Add student")
	fmt.Println("2: Edit student")
	fmt.Println("3: Delete student")
	fmt.Println("4: Return")
	fmt.Println("---------------------------")
	fmt.Println("Selection: ")

	h.sc.Scan()
	text := h.sc.Text()
	switch text {
	case "1":
		h.AddStudent()
	case "2":
		h.EditStudent()
	case "3":
		h.DeleteStudent()
	case "4":
		h.InitProgram()
	default:
		fmt.Println("Please select from 1 to 4")
	}
	h.StudentManager()
}

func (h *handler) AddStudent() {
	fmt.Println("---------------------------")
	fmt.Println("Student ID: ")
	h.sc.Scan()
	id := h.sc.Text()
	fmt.Println("Name: ")
	h.sc.Scan()
	name := h.sc.Text()
	fmt.Println("Class ID: ")
	h.sc.Scan()
	class := h.sc.Text()

	fmt.Println("---------------------------")

	student := model.Student{
		ID:      id,
		Name:    name,
		ClassID: class,
	}
	fmt.Println(student.ToString())
	if err := h.s.StudentService.Add(student); err != nil {
		panic(err)
	}
	fmt.Println("Add student success")
	h.StudentManager()
}

func (h *handler) DeleteStudent() {
	fmt.Println("---------------------------")
	fmt.Println("Student id: ")
	h.sc.Scan()
	id := h.sc.Text()

	if err := h.s.StudentService.Delete(id); err != nil {
		panic(err)
	}

	fmt.Printf("Delete student with id: %s \n", id)
	h.StudentManager()
}

func (h *handler) EditStudent() {
	fmt.Println("---------------------------")
	fmt.Println("Student id: ")
	h.sc.Scan()
	id := h.sc.Text()
	student, err := h.s.StudentService.Get(id)
	if err != nil {
		fmt.Printf("Student with ID %s not found \n", id)
		h.EditClass()
	}
	fmt.Println("Student name: ")
	h.sc.Scan()
	name := h.sc.Text()

	fmt.Println("Class ID: ")
	h.sc.Scan()
	class := h.sc.Text()

	student.Name = name
	student.ClassID = class
	if err := h.s.StudentService.Update(student); err != nil {
		panic(err)
	}

	fmt.Println("---------------------------")
	fmt.Printf("Student name: %s, class ID: %s", name, class)
	h.ClassManager()
}

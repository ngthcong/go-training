package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	repo "school-mongo/repositories"
	"school-mongo/services"
)

var (
	scanner *bufio.Scanner
	studentService *services.StudentService

)

func init(){
	ctx := context.TODO()
	db := repo.NewMongoDB()
	classRepo :=  repo.NewClassRepository(db,&ctx)
	classService := services.NewClassService(*classRepo)

	teacherRepo :=  repo.NewTeacherRepository(db,&ctx)
	teacherService := services.NewTeacherService(*teacherRepo)

	disciplineRepo :=  repo.NewDisciplineRepository(db,&ctx)
	disciplineService := services.NewDisciplineService(*disciplineRepo)

	studentRepo :=  repo.NewStudentRepository(db,&ctx)
	studentService = services.NewStudentService(*studentRepo)
}

func main() {


	scanner = bufio.NewScanner(os.Stdin)
	fmt.Println("-----School Management-----")
	fmt.Println("1: Class Management")
	fmt.Println("2: Discipline Management")
	fmt.Println("3: Teacher Management")
	fmt.Println("4: Student Management")
	fmt.Println("5: Exit")
	fmt.Println("---------------------------")
	fmt.Println("Selection: ")

	scanner.Scan()
	text := scanner.Text()

	fmt.Println(text)
	switch text {
	case "1":
		ClassManager()
	case "2":
		DisciplineManager()
	case "3":
		TeacherManager()
	case "4":
		StudentManager()
	case "5":
		os.Exit(0)
	default:
		fmt.Println("Please select from 1 to 6")
	}
	main()

}

func StudentManager() {
	fmt.Println("---------------------------")
	fmt.Println("1:Add student")
	fmt.Println("2:Edit student")
	fmt.Println("3:Delete student")
	fmt.Println("4:Return")
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
		DeleteStudent()
	case "4":
		main()
	default:
		fmt.Println("Please select from 1 to 4")
	}
	StudentManager()
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

	newStudent := &repo.Student{
		Id:      id,
		Name:    name,
		ClassID: class,
	}
	if err:=  studentService.Add(newStudent);err != nil{
		fmt.Println(err)
	}
	fmt.Println("---------------------------")
	fmt.Printf("Student name: %s Class id: %s \n", newStudent.Name, newStudent.ClassID)
	StudentManager()
}

func EditStudent() {
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

	student := &repo.Student{
		Id:      id,
		Name:    name,
		ClassID: class,
	}
	if err:=  studentService.Update(student);err != nil{
		fmt.Println(err)
	}
	fmt.Println("---------------------------")
	StudentManager()
}

func DeleteStudent() {
	fmt.Println("---------------------------")
	fmt.Println("Student id: ")
	scanner.Scan()
	id := scanner.Text()
	if err:=  studentService.Delete(id);err != nil{
		fmt.Println(err)
	}
	fmt.Println("---------------------------")
	StudentManager()
}

func TeacherManager() {
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


func DisciplineManager() {
	fmt.Println("---------------------------")
	fmt.Println("1: Add discipline")
	fmt.Println("2: Edit discipline")
	fmt.Println("3: Delete discipline")
	fmt.Println("4: Return")
	fmt.Println("---------------------------")
	fmt.Println("Selection: ")
	scanner.Scan()
	text := scanner.Text()
	switch text {
	case "1":
		AddDiscipline()
	case "2":
		EditDiscipline()
	case "3":
		DeleteDiscipline()
	case "4":
		main()
	default:
		fmt.Println("Please select from 1 to 4")
	}
}

func AddDiscipline() {
	fmt.Println("---------------------------")
	fmt.Println("Discipline id: ")
	scanner.Scan()
	id := scanner.Text()

	fmt.Println("Discipline name: ")
	scanner.Scan()
	name := scanner.Text()


	fmt.Println("---------------------------")
	fmt.Printf("Discipline name: %s \n", name)
}

func EditDiscipline() {
	fmt.Println("---------------------------")
	fmt.Println("Discipline id: ")
	scanner.Scan()
	id := scanner.Text()

	fmt.Println("Discipline name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("---------------------------")
	fmt.Printf("Teacher name: %s \n", name)
}

func DeleteDiscipline() {
	fmt.Println("---------------------------")
	fmt.Println("Discipline id: ")
	scanner.Scan()
	id := scanner.Text()
	fmt.Println("---------------------------")
	fmt.Printf("Teacher name: %s \n", id)
}

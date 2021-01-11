package main

import (
	"bufio"
	"os"
	"school-mysql/internal/handler"
	"school-mysql/internal/repositories"
	"school-mysql/internal/repositories/class"
	"school-mysql/internal/repositories/discipline"
	"school-mysql/internal/repositories/student"
	"school-mysql/internal/repositories/teacher"
	"school-mysql/internal/services"
	classSev "school-mysql/internal/services/class"
	disciplineSev "school-mysql/internal/services/discipline"
	studentSev "school-mysql/internal/services/student"
	teacherSev "school-mysql/internal/services/teacher"
)

func main() {
	db := repositories.NewDB()
	classRepo := class.New(db)
	classService := classSev.New(classRepo)

	teacherRepo := teacher.New(db)
	teacherService := teacherSev.New(teacherRepo)

	disciplineRepo := discipline.New(db)
	disciplineService := disciplineSev.New(disciplineRepo)

	studentRepo := student.New(db)
	studentService := studentSev.New(studentRepo)

	scanner := bufio.NewScanner(os.Stdin)

	service := services.New(classService, disciplineService, studentService, teacherService)

	handler := handler.New(scanner, *service)

	handler.InitProgram()
}

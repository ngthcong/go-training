package main

import (
	"bufio"
	"context"
	"os"
	handler2 "school-mongo/internal/handler"
	repo "school-mongo/internal/repositories"
	class2 "school-mongo/internal/repositories/class"
	discipline2 "school-mongo/internal/repositories/discipline"
	student2 "school-mongo/internal/repositories/student"
	teacher2 "school-mongo/internal/repositories/teacher"
	"school-mongo/internal/services"
	"school-mongo/internal/services/class"
	"school-mongo/internal/services/discipline"
	"school-mongo/internal/services/student"
	"school-mongo/internal/services/teacher"
)

func main() {
	ctx := context.TODO()
	db := repo.NewMongoDB()
	classRepo := class2.New(db, ctx)
	classService := class.New(classRepo)

	teacherRepo := teacher2.New(db, ctx)
	teacherService := teacher.New(teacherRepo)

	disciplineRepo := discipline2.New(db, ctx)
	disciplineService := discipline.New(disciplineRepo)

	studentRepo := student2.New(db, ctx)
	studentService := student.New(studentRepo)

	scanner := bufio.NewScanner(os.Stdin)

	service := services.New(classService, disciplineService, studentService, teacherService)

	handler := handler2.New(scanner, *service)

	handler.InitProgram()
}

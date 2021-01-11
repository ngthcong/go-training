package services

import (
	"school-mysql/internal/services/class"
	"school-mysql/internal/services/discipline"
	"school-mysql/internal/services/student"
	"school-mysql/internal/services/teacher"
)

type (
	Service struct {
		ClassService      class.Service
		DisciplineService discipline.Service
		StudentService    student.Service
		TeacherService    teacher.Service
	}
)

func New(classService class.Service, disciplineService discipline.Service, studentService student.Service, teacherService teacher.Service) *Service {
	return &Service{
		ClassService:      classService,
		DisciplineService: disciplineService,
		StudentService:    studentService,
		TeacherService:    teacherService,
	}
}

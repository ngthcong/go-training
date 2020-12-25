package services

import (
	"school-mongo/internal/services/class"
	"school-mongo/internal/services/discipline"
	"school-mongo/internal/services/student"
	"school-mongo/internal/services/teacher"
)

type(
	Service struct {
		classService class.Service
		disciplineService discipline.Service
		studentService student.Service
		teacherService teacher.Service
	}
)

func New(classService class.Service, disciplineService discipline.Service, studentService student.Service, teacherService teacher.Service) *Service{
	return &Service{
		classService:      classService,
		disciplineService: disciplineService,
		studentService:    studentService,
		teacherService:    teacherService,
	}
}

package services

import (
	"school-mongo/internal/services/class"
	"school-mongo/internal/services/discipline"
	"school-mongo/internal/services/student"
	"school-mongo/internal/services/teacher"
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

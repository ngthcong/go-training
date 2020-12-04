package services

import (
	"errors"
	repo "school-mongo/repositories"
)

type (
	StudentHandler interface {
		Get(id string) (*repo.Student, error)
		Add(class *repo.Student) error
		Update(class *repo.Student) error
		Delete(id string) error
	}

	StudentService struct {
		Repo repo.StudentRepository
	}
)

func NewStudentService(repo repo.StudentRepository) *StudentService {
	return &StudentService{Repo: repo}
}

func (s *StudentService) Get(id string) (student *repo.Student, err error) {
	return s.Repo.Get(id)
}

func (s *StudentService) Add(student *repo.Student) error {
	if student.Id == "" {
		return errors.New("id must not null")
	}
	if student.Name == "" {
		return errors.New("name must not null")
	}
	if student.ClassID == "" {
		return errors.New("class id must not null")
	}
	_, err := s.Repo.Get(student.Id)
	if err != nil && err.Error() != "mongo: no documents in result"{
		return err
	}
	if err := s.Repo.Add(student); err != nil {
		return err
	}
	return nil
}

func (s *StudentService) Update(student *repo.Student) error {
	_, err := s.Repo.Get(student.Id)
	if err != nil {
		return err
	}
	if err := s.Repo.Update(student); err != nil {
		return err
	}
	return nil
}

func (s *StudentService) Delete(id string) error {
	_, err := s.Repo.Get(id)
	if err != nil {
		return err
	}
	if err := s.Repo.Delete(id); err != nil {
		return err
	}
	return nil
}
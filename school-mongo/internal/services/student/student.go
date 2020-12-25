package student

import (
	"errors"
	"school-mongo/internal/model"
	"school-mongo/internal/repositories/student"
)

type (
	Service interface {
		Get(id string) (model.Student, error)
		Add(class model.Student) error
		Update(class model.Student) error
		Delete(id string) error
	}

	service struct {
		Repo student.Repository
	}
)

func New(repo student.Repository) *service {
	return &service{Repo: repo}
}

func (s *service) Get(id string) (student model.Student, err error) {
	return s.Repo.Get(id)
}

func (s *service) Add(student model.Student) error {
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

func (s *service) Update(student model.Student) error {
	_, err := s.Repo.Get(student.Id)
	if err != nil {
		return err
	}
	if err := s.Repo.Update(student); err != nil {
		return err
	}
	return nil
}

func (s *service) Delete(id string) error {
	_, err := s.Repo.Get(id)
	if err != nil {
		return err
	}
	if err := s.Repo.Delete(id); err != nil {
		return err
	}
	return nil
}
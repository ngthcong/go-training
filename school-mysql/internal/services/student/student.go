package student

import (
	"errors"
	"school-mysql/internal/model"
	"school-mysql/internal/repositories/student"
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

	v, err := s.Repo.Get(student.ID)
	if err != nil && err.Error() == model.NO_DOCUMENT {
		if err := s.Repo.Add(student); err != nil {
			return err
		}
	}
	if v.ID != "" {
		return errors.New("ID already exist")
	}
	return nil
}

func (s *service) Update(student model.Student) error {
	_, err := s.Repo.Get(student.ID)
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

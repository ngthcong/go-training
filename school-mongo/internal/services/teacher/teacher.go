package teacher

import (
	"errors"
	"school-mongo/internal/model"
	"school-mongo/internal/repositories/teacher"
)

type (
	Service interface {
		Get(id string) (model.Teacher, error)
		Add(teacher model.Teacher) error
		Update(teacher model.Teacher) error
		Delete(id string) error
	}

	service struct {
		Repo teacher.Repository
	}
)

func New(repo teacher.Repository) *service {
	return &service{Repo: repo}
}

func (s *service) Get(id string) (teacher model.Teacher, err error) {
	return s.Repo.Get(id)
}

func (s *service) Add(teacher model.Teacher) error {
	v, err := s.Repo.Get(teacher.ID)
	if err != nil && err.Error() == model.NO_DOCUMENT {
		if err := s.Repo.Add(teacher); err != nil {
			return err
		}
	}
	if v.ID != "" {
		return errors.New("ID already exist")
	}
	return nil
}

func (s *service) Update(teacher model.Teacher) error {
	_, err := s.Repo.Get(teacher.ID)
	if err != nil && err.Error() == model.NO_DOCUMENT {
		return err
	}
	if err := s.Repo.Update(teacher); err != nil {
		return err
	}
	return nil
}

func (s *service) Delete(id string) error {
	_, err := s.Repo.Get(id)
	if err != nil && err.Error() == model.NO_DOCUMENT {
		return err
	}
	if err := s.Repo.Delete(id); err != nil {
		return err
	}
	return nil
}

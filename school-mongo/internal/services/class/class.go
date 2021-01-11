package class

import (
	"errors"
	"school-mongo/internal/model"
	"school-mongo/internal/repositories/class"
)

type (
	Service interface {
		Get(id string) (model.Class, error)
		Add(class model.Class) error
		Update(class model.Class) error
		Delete(id string) error
	}

	service struct {
		Repo class.Repository
	}
)

func New(repo class.Repository) *service {
	return &service{Repo: repo}
}

func (s *service) Get(id string) (class model.Class, err error) {
	return s.Repo.Get(id)
}

func (s *service) Add(class model.Class) error {
	v, err := s.Repo.Get(class.ID)
	if err != nil && err.Error() == model.NO_DOCUMENT {
		if err := s.Repo.Add(class); err != nil {
			return err
		}
	}
	if v.ID != "" {
		return errors.New("ID already exist")
	}
	return nil
}

func (s *service) Update(class model.Class) error {
	_, err := s.Repo.Get(class.ID)
	if err != nil {
		return err
	}
	if err := s.Repo.Update(class); err != nil {
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

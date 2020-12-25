package discipline

import (
	"school-mongo/internal/model"
	"school-mongo/internal/repositories/discipline"
)

type (
	Service interface {
		Get(id string) (model.Discipline, error)
		Add(discipline model.Discipline) error
		Update(discipline model.Discipline) error
		Delete(id string) error
	}

	service struct {
		Repo discipline.Repository
	}
)

func New(repo discipline.Repository) *service {
	return &service{Repo: repo}
}

func (s *service) Get(id string) (discipline model.Discipline, err error) {
	return s.Repo.Get(id)
}

func (s *service) Add(discipline model.Discipline) error {
	_, err := s.Repo.Get(discipline.Id)
	if err != nil {
		return err
	}
	if err := s.Repo.Add(discipline); err != nil {
		return err
	}
	return nil
}

func (s *service) Update(discipline model.Discipline) error {
	_, err := s.Repo.Get(discipline.Id)
	if err != nil {
		return err
	}
	if err := s.Repo.Update(discipline); err != nil {
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

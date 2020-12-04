package services

import (
	repo "school-mongo/repositories"
)

type (
	ClassHandler interface {
		Get(id string) (*repo.Class, error)
		Add(class *repo.Class) error
		Update(class *repo.Class) error
		Delete(id string) error
	}

	ClassService struct {
		Repo repo.ClassRepository
	}
)

func NewClassService(repo repo.ClassRepository) *ClassService {
	return &ClassService{Repo: repo}
}

func (s *ClassService) Get(id string) (class *repo.Class, err error) {
	return s.Repo.Get(id)
}

func (s *ClassService) Add(class *repo.Class) error {
	_, err := s.Repo.Get(class.Id)
	if err != nil {
		return err
	}
	if err := s.Repo.Add(class); err != nil {
		return err
	}
	return nil
}

func (s *ClassService) Update(class *repo.Class) error {
	_, err := s.Repo.Get(class.Id)
	if err != nil {
		return err
	}
	if err := s.Repo.Update(class); err != nil {
		return err
	}
	return nil
}

func (s *ClassService) Delete(id string) error {
	_, err := s.Repo.Get(id)
	if err != nil {
		return err
	}
	if err := s.Repo.Delete(id); err != nil {
		return err
	}
	return nil
}

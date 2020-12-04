package services

import repo "school-mongo/repositories"

type (
	DisciplineHandler interface {
		Get(id string) (*repo.Discipline, error)
		Add(discipline *repo.Discipline) error
		Update(discipline *repo.Discipline) error
		Delete(id string) error
	}

	DisciplineService struct {
		Repo repo.DisciplineRepository
	}
)

func NewDisciplineService(repo repo.DisciplineRepository) *DisciplineService {
	return &DisciplineService{Repo: repo}
}

func (s *DisciplineService) Get(id string) (discipline *repo.Discipline, err error) {
	return s.Repo.Get(id)
}

func (s *DisciplineService) Add(discipline *repo.Discipline) error {
	_, err := s.Repo.Get(discipline.Id)
	if err != nil {
		return err
	}
	if err := s.Repo.Add(discipline); err != nil {
		return err
	}
	return nil
}

func (s *DisciplineService) Update(discipline *repo.Discipline) error {
	_, err := s.Repo.Get(discipline.Id)
	if err != nil {
		return err
	}
	if err := s.Repo.Update(discipline); err != nil {
		return err
	}
	return nil
}

func (s *DisciplineService) Delete(id string) error {
	_, err := s.Repo.Get(id)
	if err != nil {
		return err
	}
	if err := s.Repo.Delete(id); err != nil {
		return err
	}
	return nil
}

package services

import repo "school-mongo/repositories"

type (
	TeacherHandler interface {
		Get(id string) (*repo.Teacher, error)
		Add(teacher *repo.Teacher) error
		Update(teacher *repo.Teacher) error
		Delete(id string) error
	}

	TeacherService struct {
		Repo repo.TeacherRepository
	}
)

func NewTeacherService(repo repo.TeacherRepository) *TeacherService {
	return &TeacherService{Repo: repo}
}

func (s *TeacherService) Get(id string) (teacher *repo.Teacher, err error) {
	return s.Repo.Get(id)
}

func (s *TeacherService) Add(teacher *repo.Teacher) error {
	_, err := s.Repo.Get(teacher.Id)
	if err != nil {
		return err
	}
	if err := s.Repo.Add(teacher); err != nil {
		return err
	}
	return nil
}

func (s *TeacherService) Update(teacher *repo.Teacher) error {
	_, err := s.Repo.Get(teacher.Id)
	if err != nil {
		return err
	}
	if err := s.Repo.Update(teacher); err != nil {
		return err
	}
	return nil
}

func (s *TeacherService) Delete(id string) error {
	_, err := s.Repo.Get(id)
	if err != nil {
		return err
	}
	if err := s.Repo.Delete(id); err != nil {
		return err
	}
	return nil
}

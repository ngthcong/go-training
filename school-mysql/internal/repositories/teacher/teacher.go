package teacher

import (
	"database/sql"
	"school-mysql/internal/model"
)

type (
	Repository interface {
		Get(id string) (model.Teacher, error)
		Add(teacher model.Teacher) error
		Update(teacher model.Teacher) error
		Delete(id string) error
	}

	repository struct {
		db *sql.DB
	}
)

func New(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Get(id string) (teacher model.Teacher, err error) {
	err = r.db.QueryRow("SELECT * FROM teacher WHERE id=?", id).Scan(teacher.ID, teacher.Name, teacher.Disciplines)
	if err != nil {
		return teacher, err // proper error handling instead of panic in your app
	}
	return teacher, nil
}
func (r *repository) Add(teacher model.Teacher) error {
	_, err := r.db.Query("INSERT INTO teacher VALUES (?,?,?)",teacher.ID,teacher.Name,teacher.Disciplines)
	return err
}
func (r *repository) Update(teacher model.Teacher) error {
	_, err := r.db.Query("UPDATE teacher SET id=?, name=?,disciplines=?)",teacher.ID,teacher.Name,teacher.Disciplines)
	return err
}
func (r *repository) Delete(id string) error {
	_, err := r.db.Query("DELETE  FROM teacher WHERE id=?",id)
	return err
}

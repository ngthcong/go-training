package student

import (
	"database/sql"
	"school-mysql/internal/model"
)

type (
	Repository interface {
		Get(id string) (model.Student, error)
		Add(student model.Student) error
		Update(student model.Student) error
		Delete(id string) error
	}

	repository struct {
		db *sql.DB
	}
)

func New(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Get(id string) (student model.Student, err error) {
	err = r.db.QueryRow("SELECT * FROM student WHERE id=?", id).Scan(student.ID, student.Name, student.ClassID)
	if err != nil {
		return student, err // proper error handling instead of panic in your app
	}
	return student, nil
}
func (r *repository) Add(student model.Student) error {
	_, err := r.db.Query("INSERT INTO student VALUES (?,?,?)",student.ID,student.Name,student.ClassID)
	return err
}
func (r *repository) Update(student model.Student) error {
	_, err := r.db.Query("UPDATE student SET id=?, name=?,classID=?)",student.ID,student.Name,student.ClassID)
	return err
}
func (r *repository) Delete(id string) error {
	_, err := r.db.Query("DELETE  FROM student WHERE id=?",id)
	return err
}

package discipline

import (
	"database/sql"
	"school-mysql/internal/model"
)

type (
	Repository interface {
		Get(id string) (model.Discipline, error)
		Add(discipline model.Discipline) error
		Update(discipline model.Discipline) error
		Delete(id string) error
	}

	repository struct {
		db *sql.DB
	}
)

func New(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Get(id string) (discipline model.Discipline, err error) {
	err = r.db.QueryRow("SELECT * FROM discipline WHERE id=?", id).Scan(discipline.ID, discipline.Name, discipline.Lectures, discipline.Exercises)
	if err != nil {
		return discipline, err // proper error handling instead of panic in your app
	}
	return discipline, nil
}
func (r *repository) Add(discipline model.Discipline) error {
	_, err := r.db.Query("INSERT INTO discipline VALUES (?,?,?,?)", discipline.ID, discipline.Name, discipline.Lectures, discipline.Exercises)
	return err
}
func (r *repository) Update(discipline model.Discipline) error {
	_, err := r.db.Query("UPDATE discipline SET id=?, name=?,lectures=?,exercises=?)", discipline.ID, discipline.Name, discipline.Lectures, discipline.Exercises)
	return err
}
func (r *repository) Delete(id string) error {
	_, err := r.db.Query("DELETE  FROM discipline WHERE id=?", id)
	return err
}

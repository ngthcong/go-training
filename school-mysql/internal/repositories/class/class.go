package class

import (
	"database/sql"
	"school-mysql/internal/model"
)

type (
	Repository interface {
		Get(id string) (model.Class, error)
		Add(class model.Class) error
		Update(class model.Class) error
		Delete(id string) error
	}

	repository struct {
		db *sql.DB
	}
)

func New(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Get(id string) (class model.Class, err error) {
	err = r.db.QueryRow("SELECT * FROM class WHERE id=?", id).Scan(class.ID, class.Name, class.Teachers)
	if err != nil {
		return class, err // proper error handling instead of panic in your app
	}
	return class, nil
}

func (r *repository) Add(class model.Class) error {
	_, err := r.db.Query("INSERT INTO class VALUES (?,?,?)",class.ID,class.Name,class.Teachers)
	return err
}

func (r *repository) Update(class model.Class) error {
	_, err := r.db.Query("UPDATE class SET id=?, name=?,teacher=?)",class.ID,class.Name,class.Teachers)
	return err
}

func (r *repository) Delete(id string) error {
	_, err := r.db.Query("DELETE  FROM class WHERE id=?",id)
	return err
}

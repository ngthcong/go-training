package student

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"school-mongo/internal/model"
)

type (


	Repository interface {
		Get(id string) (model.Student, error)
		Add(student model.Student) error
		Update(student model.Student) error
		Delete(id string) error
	}

	repository struct {
		col *mongo.Collection
		ctx context.Context
	}
)

func New(db *mongo.Database, ctx context.Context) *repository {
	c := db.Collection("student")
	return &repository{col: c, ctx: ctx}
}

func (r *repository) Get(id string) (student model.Student, err error) {
	err = r.col.FindOne(r.ctx, bson.D{{"id", id}}).Decode(&student)
	return
}
func (r *repository) Add(student model.Student) error {
	_, err := r.col.InsertOne(r.ctx, student)
	return err
}
func (r *repository) Update(student model.Student) error {
	filter := bson.D{{"id", student.Id}}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{ReturnDocument: &after}
	update := bson.D{{"$set", bson.D{{"name", student.Name},
		{"class_id", student.ClassID}}}}
	err := r.col.FindOneAndUpdate(r.ctx, filter, update, &returnOpt)
	return err.Err()
}
func (r *repository) Delete(id string) error {
	err := r.col.FindOneAndDelete(r.ctx, bson.D{{"id", id}})
	return err.Err()
}

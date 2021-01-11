package teacher

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"school-mongo/internal/model"
)

type (
	Repository interface {
		Get(id string) (model.Teacher, error)
		Add(teacher model.Teacher) error
		Update(teacher model.Teacher) error
		Delete(id string) error
	}

	repository struct {
		col *mongo.Collection
		ctx context.Context
	}
)

func New(db *mongo.Database, ctx context.Context) *repository {
	c := db.Collection("teacher")
	return &repository{col: c, ctx: ctx}
}

func (r *repository) Get(id string) (teacher model.Teacher, err error) {
	err = r.col.FindOne(r.ctx, bson.D{{"id", id}}).Decode(&teacher)
	return
}
func (r *repository) Add(teacher model.Teacher) error {
	_, err := r.col.InsertOne(r.ctx, teacher)
	return err
}
func (r *repository) Update(teacher model.Teacher) error {
	filter := bson.D{{"id", teacher.ID}}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{ReturnDocument: &after}
	update := bson.D{{"$set", bson.D{{"name", teacher.Name}}}}
	err := r.col.FindOneAndUpdate(r.ctx, filter, update, &returnOpt)
	return err.Err()
}
func (r *repository) Delete(id string) error {
	err := r.col.FindOneAndDelete(r.ctx, bson.D{{"id", id}})
	return err.Err()
}

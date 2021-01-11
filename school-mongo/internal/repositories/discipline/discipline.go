package discipline

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"school-mongo/internal/model"
)

type (
	Repository interface {
		Get(id string) (model.Discipline, error)
		Add(discipline model.Discipline) error
		Update(discipline model.Discipline) error
		Delete(id string) error
	}

	repository struct {
		col *mongo.Collection
		ctx context.Context
	}
)

func New(db *mongo.Database, ctx context.Context) *repository {
	c := db.Collection("discipline")
	return &repository{col: c, ctx: ctx}
}

func (r *repository) Get(id string) (discipline model.Discipline, err error) {
	err = r.col.FindOne(r.ctx, bson.D{{"id", id}}).Decode(&discipline)
	return
}
func (r *repository) Add(discipline model.Discipline) error {
	_, err := r.col.InsertOne(r.ctx, discipline)
	return err
}
func (r *repository) Update(discipline model.Discipline) error {
	filter := bson.D{{"id", discipline.ID}}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{ReturnDocument: &after}
	update := bson.D{{"$set", bson.D{{"name", discipline.Name},
		{"lectures", discipline.Lectures}, {"exercises", discipline.Exercises}}}}
	err := r.col.FindOneAndUpdate(r.ctx, filter, update, &returnOpt)
	return err.Err()
}
func (r *repository) Delete(id string) error {
	err := r.col.FindOneAndDelete(r.ctx, bson.D{{"id", id}})
	return err.Err()
}

package class

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"school-mongo/internal/model"
)

type (
	Repository interface {
		Get(id string) (model.Class, error)
		Add(class model.Class) error
		Update(class model.Class) error
		Delete(id string) error
	}

	repository struct {
		col *mongo.Collection
		ctx context.Context
	}
)

func New(db *mongo.Database, ctx context.Context) *repository {
	c := db.Collection("class")
	return &repository{col: c, ctx: ctx}
}

func (r *repository) Get(id string) (class model.Class, err error) {
	err = r.col.FindOne(r.ctx, bson.D{{"id", id}}).Decode(&class)
	return
}

func (r *repository) Add(class model.Class) error {
	_, err := r.col.InsertOne(r.ctx, class)
	return err
}

func (r *repository) Update(class model.Class) error {
	filter := bson.D{{"id", class.ID}}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{ReturnDocument: &after}
	update := bson.D{{"$set", bson.D{{"name", class.Name}}}}
	err := r.col.FindOneAndUpdate(r.ctx, filter, update, &returnOpt)
	return err.Err()
}

func (r *repository) Delete(id string) error {
	err := r.col.FindOneAndDelete(r.ctx, bson.D{{"id", id}})
	return err.Err()
}

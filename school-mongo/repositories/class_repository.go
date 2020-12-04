package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Class struct {
	Id      string `bson:"id"`
	Name    string `bson:"name"`
}

type ClassHandler interface {
	Get(id string) (*Class, error)
	Add(class *Class) error
	Update(class *Class) error
	Delete(id string) error
}

type ClassRepository struct {
	col *mongo.Collection
	ctx *context.Context
}

func NewClassRepository(db *mongo.Database) *DisciplineRepository {
	c := db.Collection("discipline")
	return &DisciplineRepository{col: c}
}

func (r *ClassRepository) Get(id string) (class *Class, err error) {
	err = r.col.FindOne(*r.ctx, bson.D{{"id", id}}).Decode(&class)
	return
}

func (r *ClassRepository) Add(class *Class) error {
	_, err := r.col.InsertOne(*r.ctx, class)
	return err
}

func (r *ClassRepository) Update(class *Class) error {
	filter := bson.D{{"id", class.Id}}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{ReturnDocument: &after}
	update := bson.D{{"$set", bson.D{{"name", class.Name}}}}
	err := r.col.FindOneAndUpdate(ctx, filter, update, &returnOpt)
	return err.Err()
}

func (r *ClassRepository) Delete(id string) error {
	err := r.col.FindOneAndDelete(*r.ctx, bson.D{{"id", id}})
	return err.Err()
}

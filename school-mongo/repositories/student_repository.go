package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student struct {
	Id      string `bson:"id"`
	Name    string `bson:"name"`
	ClassID string `bson:"class_id"`
}

type StudentHandler interface {
	Get(id string) (*Student, error)
	Add(student *Student) error
	Update(student *Student) error
	Delete(id string) error
}

type StudentRepository struct {
	col *mongo.Collection
	ctx *context.Context
}

func NewStudentRepository(db *mongo.Database) *StudentRepository {
	c := db.Collection("student")
	return &StudentRepository{col: c}
}

func (r *StudentRepository) Get(id string) (student *Student, err error) {
	err = r.col.FindOne(*r.ctx, bson.D{{"id", id}}).Decode(&student)
	return
}
func (r *StudentRepository) Add(student *Student) error {
	_, err := r.col.InsertOne(*r.ctx, student)
	return err
}
func (r *StudentRepository) Update(student *Student) error {
	filter := bson.D{{"id", student.Id}}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{ReturnDocument: &after}
	update := bson.D{{"$set", bson.D{{"name", student.Name},
		{"class_id", student.ClassID}}}}
	err := r.col.FindOneAndUpdate(ctx, filter, update, &returnOpt)
	return err.Err()
}
func (r *StudentRepository) Delete(id string) error {
	 err := r.col.FindOneAndDelete(*r.ctx,  bson.D{{"id",id }})
	return err.Err()
}

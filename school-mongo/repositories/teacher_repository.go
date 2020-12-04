package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Teacher struct {
	Id      string `bson:"id"`
	Name    string `bson:"name"`
	DisciplineID string `bson:"discipline_id"`
}

type TeacherHandler interface {
	Get(id string) (*Teacher, error)
	Add(teacher *Teacher) error
	Update(teacher *Teacher) error
	Delete(id string) error
}

type TeacherRepository struct {
	col *mongo.Collection
	ctx *context.Context
}

func NewTeacherRepository(db *mongo.Database) *TeacherRepository {
	c := db.Collection("teacher")
	return &TeacherRepository{col: c}
}

func (r *TeacherRepository) Get(id string) (student *Student, err error) {
	err = r.col.FindOne(*r.ctx, bson.D{{"id", id}}).Decode(&student)
	return
}
func (r *TeacherRepository) Add(student *Student) error {
	_, err := r.col.InsertOne(*r.ctx, student)
	return err
}
func (r *TeacherRepository) Update(student *Student) error {
	filter := bson.D{{"id", student.Id}}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{ReturnDocument: &after}
	update := bson.D{{"$set", bson.D{{"name", student.Name},
		{"class_id", student.ClassID}}}}
	err := r.col.FindOneAndUpdate(ctx, filter, update, &returnOpt)
	return err.Err()
}
func (r *TeacherRepository) Delete(id string) error {
	err := r.col.FindOneAndDelete(*r.ctx,  bson.D{{"id",id }})
	return err.Err()
}

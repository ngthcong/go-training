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

func NewTeacherRepository(db *mongo.Database,ctx *context.Context) *TeacherRepository {
	c := db.Collection("teacher")
	return &TeacherRepository{col: c,ctx: ctx}
}

func (r *TeacherRepository) Get(id string) (teacher *Teacher, err error) {
	err = r.col.FindOne(*r.ctx, bson.D{{"id", id}}).Decode(&teacher)
	return
}
func (r *TeacherRepository) Add(teacher *Teacher) error {
	_, err := r.col.InsertOne(*r.ctx, teacher)
	return err
}
func (r *TeacherRepository) Update(teacher *Teacher) error {
	filter := bson.D{{"id", teacher.Id}}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{ReturnDocument: &after}
	update := bson.D{{"$set", bson.D{{"name", teacher.Name}}}}
	err := r.col.FindOneAndUpdate(ctx, filter, update, &returnOpt)
	return err.Err()
}
func (r *TeacherRepository) Delete(id string) error {
	err := r.col.FindOneAndDelete(*r.ctx,  bson.D{{"id",id }})
	return err.Err()
}

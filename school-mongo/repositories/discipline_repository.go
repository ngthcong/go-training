package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Discipline struct {
	Id        string `bson:"id"`
	Name      string `bson:"name"`
	Lectures  int    `bson:"lectures"`
	Exercises int    `bson:"exercises"`
}

type DisciplineHandler interface {
	Get(id string) (*Discipline, error)
	Add(Discipline *Discipline) error
	Update(Discipline *Discipline) error
	Delete(id string) error
}

type DisciplineRepository struct {
	col *mongo.Collection
	ctx *context.Context
}

func NewDisciplineRepository(db *mongo.Database) *DisciplineRepository {
	c := db.Collection("discipline")
	return &DisciplineRepository{col: c}
}

func (r *DisciplineRepository) Get(id string) (discipline *Discipline, err error) {
	err = r.col.FindOne(*r.ctx, bson.D{{"id", id}}).Decode(&discipline)
	return
}
func (r *DisciplineRepository) Add(discipline *Discipline) error {
	_, err := r.col.InsertOne(*r.ctx, discipline)
	return err
}
func (r *DisciplineRepository) Update(discipline *Discipline) error {
	filter := bson.D{{"id", discipline.Id}}
	after := options.After
	returnOpt := options.FindOneAndUpdateOptions{ReturnDocument: &after}
	update := bson.D{{"$set", bson.D{{"name", discipline.Name},
		{"lectures", discipline.Lectures},{"exercises",discipline.Exercises}}}}
	err := r.col.FindOneAndUpdate(ctx, filter, update, &returnOpt)
	return err.Err()
}
func (r *DisciplineRepository) Delete(id string) error {
	err := r.col.FindOneAndDelete(*r.ctx, bson.D{{"id", id}})
	return err.Err()
}

package model

import "fmt"

type (
	Class struct {
		Id   string `bson:"id"`
		Name string `bson:"name"`
	}
	Discipline struct {
		Id        string `bson:"id"`
		Name      string `bson:"name"`
		Lectures  int    `bson:"lectures"`
		Exercises int    `bson:"exercises"`
	}
	Student struct {
		Id      string `bson:"id"`
		Name    string `bson:"name"`
		ClassID string `bson:"class_id"`
	}
	Teacher struct {
		Id          string `bson:"id"`
		Name        string `bson:"name"`
		ClassID     string `bson:"class_id"`
		Disciplines string `bson:"disciplines"`
	}
)

func (t Teacher) ToString() string {
	return fmt.Sprintf("Teacher ID: %s, name: %s, discipline ID: %s", t.Id, t.Name, t.DisciplineID)
}
func (t Student) ToString() string {
	return fmt.Sprintf("Student ID: %s, name: %s, class ID: %s", t.Id, t.Name, t.ClassID)
}
func (t Discipline) ToString() string {
	return fmt.Sprintf("Discipline ID: %s, name: %s, Lecture : %d, Exercise: %d", t.Id, t.Name, t.Lectures, t.Exercises)
}
func (t Class) ToString() string {
	return fmt.Sprintf("Class ID: %s, name: %s", t.Id, t.Name)
}

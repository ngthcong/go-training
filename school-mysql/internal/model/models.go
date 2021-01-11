package model

import (
	"fmt"
)

const (
	NO_DOCUMENT = "mongo: no documents in result"
)

//Class ....
type Class struct {
	ID       string `bson:"id"`
	Name     string `bson:"name"`
	Teachers string `bson:"teachers"`
}

//Discipline ....
type Discipline struct {
	ID        string `bson:"id"`
	Name      string `bson:"name"`
	Lectures  string `bson:"lectures"`
	Exercises string `bson:"exercises"`
}

//Student ....
type Student struct {
	ID      string `bson:"id"`
	Name    string `bson:"name"`
	ClassID string `bson:"class_id"`
}

//Teacher ....
type Teacher struct {
	ID          string `bson:"id"`
	Name        string `bson:"name"`
	Disciplines string `bson:"disciplines"`
}

//ToString ...
func (t Teacher) ToString() string {
	return fmt.Sprintf("Teacher ID: %s, name: %s, discipline ID: %s", t.ID, t.Name, t.Disciplines)
}

//ToString ...
func (t Student) ToString() string {
	return fmt.Sprintf("Student ID: %s, name: %s, class ID: %s", t.ID, t.Name, t.ClassID)
}

//ToString ...
func (t Discipline) ToString() string {
	return fmt.Sprintf("Discipline ID: %s, name: %s, Lecture : %d, Exercise: %d", t.ID, t.Name, t.Lectures, t.Exercises)
}

//ToString ...
func (t Class) ToString() string {
	return fmt.Sprintf("Class ID: %s, name: %s", t.ID, t.Name)
}

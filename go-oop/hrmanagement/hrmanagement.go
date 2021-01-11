package main

import (
	"fmt"
	"sort"
)

type (
	Human struct {
		Name     string
		LastName string
	}
	Student struct {
		human Human
		Grade int
	}
	Worker struct {
		human           Human
		WeekSalary      float64
		WorkHoursPerDay float64
	}

	Humans struct {
		list []IHuman
	}
	IStudent interface {
		SortByGrade(desc bool)
		ToString()
	}
	IWorker interface {
		SortByMoney(desc bool)
	}

	IHuman interface {
		ToName() string
		ToLastName() string
	}

	IHumans interface {
		SortByName(desc bool)
		SortByLastName(desc bool)
	}

	Students struct {
		list []Student
	}

	Workers struct {
		list []Worker
	}
)
func(s Student) ToName() string {
	return s.human.Name
}

func(w Worker) ToName() string {
	return w.human.Name
}

func(s Student) ToLastName() string {
	return s.human.LastName
}

func(w Worker) ToLastName() string {
	return w.human.LastName
}


func (w Worker) MoneyPerHour() float64 {
	return (w.WeekSalary / 5) / w.WorkHoursPerDay
}

func (s *Students) SortByGrade(desc bool) {
	if desc {
		sort.SliceStable(s.list, func(i, j int) bool {
			return s.list[i].Grade > s.list[j].Grade
		})
	}
	sort.SliceStable(s.list, func(i, j int) bool {
		return s.list[i].Grade < s.list[j].Grade
	})
}
func (w *Workers) SortByMoney(desc bool) {
	if desc {
		sort.SliceStable(w.list, func(i, j int) bool {
			return w.list[i].MoneyPerHour() > w.list[j].MoneyPerHour()
		})
	}
	sort.SliceStable(w.list, func(i, j int) bool {
		return w.list[i].MoneyPerHour() < w.list[j].MoneyPerHour()
	})
}

func(h Humans) SortByName(desc bool){
	if desc {
		sort.SliceStable(h.list, func(i, j int) bool {
			return h.list[i].ToName() > h.list[j].ToName()
		})
	}
	sort.SliceStable(h.list, func(i, j int) bool {
		return h.list[i].ToName() < h.list[j].ToName()
	})
}
func(h Humans) SortByLastName(desc bool){
	if desc {
		sort.SliceStable(h.list, func(i, j int) bool {
			return h.list[i].ToLastName() > h.list[j].ToLastName()
		})
	}
	sort.SliceStable(h.list, func(i, j int) bool {
		return h.list[i].ToLastName() < h.list[j].ToLastName()
	})
}



func (s Students) ToString() {
	for i, v := range s.list {
		fmt.Println("")
		fmt.Printf("#%d \n", i)
		fmt.Printf("Student first name: %s \n", v.human.Name)
		fmt.Printf("Student last name: %s \n", v.human.LastName)
		fmt.Printf("Grade: %v \n", v.Grade)
		fmt.Println("")
	}
}
func (w Workers) ToString() {
	for i, v := range w.list {
		fmt.Println("")
		fmt.Printf("#%d \n", i)
		fmt.Printf("Worker first name: %s \n", v.human.Name)
		fmt.Printf("Worker last name: %s \n", v.human.LastName)
		fmt.Printf("Week salary: %f \n", v.WeekSalary)
		fmt.Printf("Work hour per day: %f \n", v.WorkHoursPerDay)
		fmt.Printf("Salary per hour: %f \n", v.MoneyPerHour())
		fmt.Println("")
	}
}
func (h Humans) ToString() {

	for i, v := range h.list {
		fmt.Println("")
		fmt.Printf("#%d \n", i)
		fmt.Printf("Name: %s \n", v.ToName())
		fmt.Printf("Last name: %s \n", v.ToLastName())
		fmt.Println("")
	}

}
func main() {


	studentList := []Student{
		{
			human: Human{
				Name:     "Long",
				LastName: "Nguyen",
			},
			Grade: 9,
		},
		{
			human: Human{
				Name:     "Binh",
				LastName: "Le",
			},
			Grade: 7,
		},
		{
			human: Human{
				Name:     "Lan",
				LastName: "Phong",
			},
			Grade: 8,
		},
	}

	students := Students{list: studentList}

	workerList := []Worker{
		{
			human: Human{
				Name:     "Cong",
				LastName: "Lam",
			},
			WeekSalary: 900,
			WorkHoursPerDay: 8,
		},
		{
			human: Human{
				Name:     "Tuon",
				LastName: "Ha",
			},
			WeekSalary: 500,
			WorkHoursPerDay: 5,
		},
		{
			human: Human{
				Name:     "Quyen",
				LastName: "Binh",
			},
			WeekSalary: 1000,
			WorkHoursPerDay: 8,
		},
	}

	workers := Workers{list:workerList }
 	var humans []IHuman
	for _,v := range studentList{
		humans = append(humans, v)
	}
	for _,v := range workerList{
		humans = append(humans, v)
	}
fmt.Println("Student sort by grade")
	students.SortByGrade(true)
	students.ToString()
	fmt.Println("Worker sort by salary")
	workers.SortByMoney(true)
	workers.ToString()


	humanList := Humans{list: humans}
	fmt.Println("Human sort by Name")
	humanList.SortByName(true)
	humanList.ToString()
	fmt.Println("Human sort by LastName")
	humanList.SortByLastName(true)
	humanList.ToString()
}

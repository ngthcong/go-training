package main

import "fmt"

type (
	IAnimal interface {
		ToSound() string
		Gender() string
		ToAge() int
		ToName() string
	}
	Cat struct {
		Age   int
		Name  string
		Sex   string
		Sound string
	}
	Dog struct {
		Age   int
		Name  string
		Sex   string
		Sound string
	}
	Frog struct {
		Age   int
		Name  string
		Sex   string
		Sound string
	}
	Kitten struct {
		Cat Cat
	}
	Tomcat struct {
		Cat Cat
	}
	Animals struct {
		Specie string
		list   []IAnimal
	}
)

func (a Animals) AverageAge() {
	var total int
	for _, v := range a.list {
		total += v.ToAge()
	}
	result := float64(total) / float64(len(a.list))
	fmt.Printf("Avarage age of %v is: %v \n", a.Specie, result)
}

func (a Animals) ToString() {
	for i, v := range a.list {
		fmt.Println("")
		fmt.Printf("#%d \n", i)
		fmt.Printf("Animal Specie: %s \n", a.Specie)

		fmt.Printf("Animal name: %s \n", v.ToName())
		fmt.Printf("Sex: %v \n", v.Gender())
		fmt.Printf("Age: %v \n", v.ToAge())
		fmt.Printf("Sound: %s \n", v.ToSound())
		fmt.Println("")
	}
}

func (c Cat) ToName() string {
	return c.Name
}
func (d Dog) ToName() string {
	return d.Name
}
func (f Frog) ToName() string {
	return f.Name
}
func (k Kitten) ToName() string {
	return k.Cat.ToName()
}
func (t Tomcat) ToName() string {
	return t.Cat.ToName()
}

func (c Cat) ToSound() string {
	return c.Sound
}
func (f Frog) ToSound() string {
	return f.Sound
}
func (d Dog) ToSound() string {
	return d.Sound
}
func (k Kitten) ToSound() string {
	return k.Cat.ToSound()
}
func (t Tomcat) ToSound() string {
	return t.Cat.ToSound()
}

func (c Cat) Gender() string {
	return c.Sex
}
func (f Frog) Gender() string {
	return f.Sex
}
func (d Dog) Gender() string {
	return d.Sex
}
func (k Kitten) Gender() string {
	return "Female"
}
func (t Tomcat) Gender() string {
	return "Male"
}

func (c Cat) ToAge() int {
	return c.Age
}
func (f Frog) ToAge() int {
	return f.Age
}
func (d Dog) ToAge() int {
	return d.Age
}

func (k Kitten) ToAge() int {
	return k.Cat.ToAge()
}
func (t Tomcat) ToAge() int {
	return t.Cat.ToAge()
}

func main() {
	dog1 := Dog{Name: "Lulu", Age: 2, Sex: "Male", Sound: "Gau"}
	dog2 := Dog{Name: "Sam", Age: 3, Sex: "Female", Sound: "Gauuu"}
	dog3 := Dog{Name: "Ali", Age: 5, Sex: "Male", Sound: "Gauuuuuu"}
	cat1 := Cat{Name: "Ginger", Age: 5, Sex: "Female", Sound: "Meo"}
	cat2 := Cat{Name: "Brade", Age: 3, Sex: "Male", Sound: "Meow"}
	cat3 := Cat{Name: "Nana", Age: 1, Sex: "Female", Sound: "Maowww"}

	kitten1 := Kitten{Cat: Cat{
		Age:   1,
		Name:  "Nunu",
		Sound: "Miu",
	}}
	tomcat1 := Tomcat{Cat: Cat{
		Age:   2,
		Name:  "Lala",
		Sound: "Miuuuu",
	}}

	frog1 := Frog{Name: "Nana", Age: 1, Sex: "Female", Sound: "Purreekaaaa"}
	frog2 := Frog{Name: "Xi", Age: 2, Sex: "Male", Sound: "Purreekuuuu"}
	frog3 := Frog{Name: "Cu", Age: 4, Sex: "Female", Sound: "Purreekiii"}
	dogs := Animals{
		Specie: "Dog",
		list:   []IAnimal{dog1, dog2, dog3},
	}
	cats := Animals{
		Specie: "Cat",
		list:   []IAnimal{cat1, cat2, cat3, kitten1, tomcat1},
	}
	frogs := Animals{
		Specie: "Frog",
		list:   []IAnimal{frog1, frog2, frog3},
	}

	cats.ToString()
	dogs.ToString()
	frogs.ToString()
	dogs.AverageAge()
	cats.AverageAge()
	frogs.AverageAge()
}

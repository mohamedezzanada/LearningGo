package main

import (
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/mohamedezzatnada/GolangProject/helpers"
)

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Camel struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

var s = "seven"

type User struct {
	FirstName string
	LastName  string
	PhoneName string
	Age       int
	BirthDate time.Time
}

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	log.Println("s is", s)
	var s2 = "six"
	log.Println("s2 is", s2)
	var whatToSay string
	var myString string
	myString = "Green"
	whatToSay = saySomething("Hello")
	var i int
	i = 7
	i = 8
	log.Printf(whatToSay)
	fmt.Println(whatToSay)
	log.Println(i)
	log.Println("myString is set to", myString)
	changeUsingPointer(&myString)
	log.Println("after func call myString is set to", myString)

	user := User{
		FirstName: "Mohamed",
		LastName:  "Nada",
		PhoneName: "00000000",
	}
	log.Println(user.FirstName, user.LastName, user.BirthDate)

	var myVar myStruct
	myVar.FirstName = "Mohamed"
	myVar2 := myStruct{
		FirstName: "Nada",
	}

	log.Println("myVar is set to", myVar.FirstName)
	log.Println("myVar2 is set to", myVar2.FirstName)

	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())

	myMap := make(map[string]string)
	myMap["TheKey"] = "TheValue"
	myMap["TheOtherKey"] = "TheOtherValue"

	log.Println(myMap["TheKey"])
	log.Println(myMap["TheOtherKey"])

	myMap1 := make(map[string]User)

	me := User{
		FirstName: "Mohamed",
		LastName:  "Nada",
	}

	myMap1["me"] = me
	log.Println(myMap1["me"].FirstName)

	var myNewVar float32
	myNewVar = 11.1

	log.Println(myNewVar)

	var mySlice []string
	mySlice = append(mySlice, "Mohamed")
	mySlice = append(mySlice, "Nada")

	log.Println(mySlice)

	var myNewSlice []int
	myNewSlice = append(myNewSlice, 2)
	myNewSlice = append(myNewSlice, 1)
	myNewSlice = append(myNewSlice, 3)

	sort.Ints(myNewSlice)

	log.Println(myNewSlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers)
	log.Println(numbers[0:2])

	names := []string{"one", "two", "fish", "cat"}

	log.Println(names)

	isTrue := false

	if isTrue == true {
		log.Println("isTrue", isTrue)
	} else {
		log.Println("isTrue", isTrue)
	}

	myNum := 100
	isTrue1 := false

	if myNum > 99 && !isTrue1 {
		log.Println("myNum is greater than 99 and isTrue1 is set to true")
	} else if myNum == 101 && !isTrue1 {
		log.Println("1")

	} else if myNum > 1000 && !isTrue1 {
		log.Println("2")
	}

	switch myNum {
	case 1:
		log.Println("100")
	case 2:
		log.Println("200")
	case 3:
		log.Println("300")
	default:
		log.Println("Something else")

	}

	for i := 0; i <= 10; i++ {
		log.Println(i)
	}
	myNewSlice1 := []string{"one", "two", "three", "four"}

	for j, x := range myNewSlice1 {
		log.Println(x, j)
	}

	myNewMap := make(map[string]string)
	myNewMap["FirstValue"] = "one"
	myNewMap["SecondValue"] = "two"
	myNewMap["ThirdValue"] = "three"

	for i, x := range myNewMap {
		log.Println(i, x)
	}

	var userSlice []User

	u1 := User{
		FirstName: "Mohamed",
	}

	u2 := User{
		FirstName: "Ahmed",
	}

	userSlice = append(userSlice, u1)
	userSlice = append(userSlice, u2)

	for i, x := range userSlice {
		log.Println(i, x.FirstName)
	}

	camel := Camel{
		Name:  "Camel Name",
		Breed: "Breeder Name",
	}

	gorilla := Gorilla{
		Name:          "King Kong",
		Color:         "Black",
		NumberOfTeeth: 32,
	}

	PrintInfo(camel)
	PrintInfo(gorilla)
	var newVar helpers.SomeType
	newVar.TypeName = "Mohamed Type"
	log.Println(newVar.TypeName)

}

func (c Camel) Says() string {
	return "return from camel function"
}

func (c Camel) NumberOfLegs() int {
	return 4
}

func (g Gorilla) Says() string {
	return "return from gorilla function"
}

func (g Gorilla) NumberOfLegs() int {
	return 2
}

func PrintInfo(a Animal) {
	log.Println("This animal says", a.Says(), " and has ", a.NumberOfLegs(), "legs")
}

func saySomething(s string) string {
	return s
}
func saySomethingNew(s string) (string, string) {
	return s, "World"
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}

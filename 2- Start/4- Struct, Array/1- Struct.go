package ___Struct__Array

import "fmt"

/*

# Struct data type - tool:
in programming, we are always modeling ==> the world around us, our ideas & etc...

Ex: a software for university, school, hospital, ATM, Gas station and...

in university software, we have different kinds of entities.
Ex: student; class; lesson; teacher and...

each entity/class has its own properties and methods
but during modeling; We are not going to use all the properties and methods in our model.
according to our software requirements, we choose which properties and methods to have.

Ex: in university software there is no need to data of student's eye color to be a property :D

*/

// we can use struct tool, to create a new type - a new model:

// basic types   ==> int, string, bool & etc...
// compound type ==> struct

func main() {

	// type {name of the entity which you want to model} struct
	// Ex: type Student struct ==> this is a model of Student entity base on the requirements of our software

	type Student struct {
		ID           uint
		NationNumber string
		Age          int
		FirstName    string
		isActive     bool
		// this is a private property, so it's not access able from another package - module :D
		// we can access it only in current file (1- Struct.go) and you can give it value.
	}

	// how to use this struct?
	// remember ==> when we wanted to create a string variable we did like this:
	var firstName string
	firstName = "amir"
	fmt.Println(firstName)

	// to use struct:
	var u Student
	// this 'u' variable type is Student not int, string, bool or...
	// all properties - fields of this struct are Zero Value ==> the integers are 0 and strings are empty string

	// to access the Student properties:
	fmt.Println(u.Age)
	fmt.Println(u.NationNumber)
	fmt.Println(u.FirstName)

	// =======================================

	// initializing the struct:

	var i = Student{
		ID:           1,
		NationNumber: "434433",
		Age:          31,
		FirstName:    "test",
	}

	// if you don't give value to a property of a struct, it will be zero value :D

	fmt.Println(i.Age)
	fmt.Println(i.NationNumber)
	fmt.Println(i.FirstName)
	fmt.Println(i.ID)

	// =======================================
	// another way to initialize the struct:

	var t Student
	t.FirstName = "amir"
	t.Age = 31

	// if you don't give value to a property of a struct, it will be zero value :D

	// =======================================

	// another way to initialize the struct:
	// here, i want to initialize the struct without defining the name of the property
	// in this way, we have to obey the order

	var j = Student{1, "142", 4, "amir", true}
	fmt.Println(j)

	// =========================================

	fmt.Println("My name is", j.FirstName, "My age is", j.Age)

	// use Printf instead of above,  which accepts formatter:
	// each formater is used for a data type and we have to pass that type of data type
	// Ex: %s ==> for string | %d ==> for integer

	fmt.Printf("My name is %s my age is %d \n", j.FirstName, j.Age)

}

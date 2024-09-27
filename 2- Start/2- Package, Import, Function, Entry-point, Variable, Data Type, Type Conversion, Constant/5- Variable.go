package _2

import "fmt"

var first_name, last_name string
var age int

// you have to define the type of the variable.

// a box (bottle) for holding acid
// a box for holding toilet paper

// base on the type of the things we want to store,
// we should have different container - holders for them

// it's same for variable bro :D ==> you should define the type of the thing
// you want to store, then choose a container for that :D

// Ex: we cant store liquid in a paper base container

// ===================================

// you can even define the primitive value
var i, j = 1, 2
var f, j1 = true, true

// when you give primitive value to a variable, you can even not define the data type,
// because compiler can predict the data type of the primitive data

var cc, bb bool = true, false

var cfd int = 100

// ====================================

// for giving primitive data type there is a simple way:

func main() {

	points := 100
	fmt.Println(points)

	first, last, age1 := "amir", "tvk", 31
	fmt.Println(first, last, age1)
	// compiler will convert points to ==> var points int = 100

}

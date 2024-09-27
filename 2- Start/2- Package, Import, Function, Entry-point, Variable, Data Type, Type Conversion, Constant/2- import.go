package _2

/*
Imagine we have a module for working with files, another module for working with database
and other modules...

The question is, how we can use these modules in our project? ==> use Import tool.

Import means, using/importing one or several functionalities/methods of a package in your codes.

in below code, we have used/imported the functionality/method of generating a random number &
the functionality/method of printing a text in output.
*/

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Print("my favorite number is ", rand.Intn(100))
}

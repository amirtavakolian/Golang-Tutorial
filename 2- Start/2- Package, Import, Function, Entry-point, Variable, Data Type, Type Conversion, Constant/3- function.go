package _2

import "fmt"

func calculateBasketPrice(x int, y int, z int) int {
	return x + y + z
}

/*

What is Function?
When are code base get bigger & bigger, maybe we duplicate several logics/algorithms.

Ex:
An algorithm/logic, for calculating the total price of the basket,
which we need it in different parts of our program.

so one way is to duplicate this logic in several parts in our software.
so if you want to make changes in the logic we have to change it in all the places
==>
It's very time-consuming and leads to making bugs

-------------------------------------------

The second way is to use function tool ==> it's a good tool for making our codes reuse-able.
so, we can use that code in several places in our software

We call it reuse-ability ==> just call the function.

Ex:
a function for calculating the basket price and call it from any-where you want.
and if you want to make changes in the algorithm of calculating basket just make changes
in one place.

-------------------------------------------

# Data type of variables and parameters:

Each function can get data from input(parameters) and make process on those data and return the result.
Example of input can be the console which user can enter inputs and the program process them and return the result.

The functions also can get input and process them and return something

Ex of input ==> fmt.Print("hello there") ==> hello there is an input for print function

We have to define the type of parameters (inputs of the function) and the type of return type
of that function ==> integer, string, or...

golang is a type safe language which means you cannot define a variable or parameter without defining its type

*/

func calculateDiscount(x int, y int, z int) int {
	return x + y + z
}

// if all the parameters type are same, just put one type at the end:
func calculateVAT(x, y int) int {
	return x + y
}

// -------------------------------------------------------------

// you can even have several outputs:

func calculateBasket(x, y int) (int, int) {
	return x + y, y - x
}

func main() {

	first, second := calculateBasket(10, 50)
	fmt.Println(first, second)
}

// -------------------------------------------------------------

// if you define name for the outputs, just use return
// there is no need to do like above example ==> return x + y, y - x

func calculate2(x, y int) (x1, y2 int) {
	x1 = x * y
	y2 = y * 2 * x
	return
}

// -------------------------------------------------------------

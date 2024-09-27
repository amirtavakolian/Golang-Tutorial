package _2

import "fmt"

/*
in golang we have basic types boolean(true, false), integer
(int, int8, int16, int32, int64, uint, unit8, uint16 &...),  string, float, map, rune &...

We have different integer ==> the difference between them is in the size of the value that they can accept

int8 ==> 2- Package, Import, Function, Entry-point, Variable, Data Type, Type Conversion, Constant pow 7 ==> -128 till +127

imagine you want to create carton(container) for different things.
We have to create the cartons based on the size of the things that we want to store
We don't use a carton in size 400x500 for a thing with size 70x70  ==> Ex: carton of shoes

So if our value is not more than 128, we have to use int8 ==> a 75x75 carton for a thing with size of 70x70
size of container  = size of the thing - value we want to store

uint ==> unsigned int ==> 0 till 255
uint8 ==> unsigned int 8

================================================================================

# Zero Value:

In other languages with define variables & we forget to give them value and in the code
we use that variable for divisions, sum, minus &...

So our program will face to bugs and exception or error happen


golang has a rule for this:
when you declare a variable, you even give it a value or the compiler will give the variables default value
So the variables will not be without value in golang


# this action of compiler is known as zero value

Ex:
The zero value for integer 0
The zero value for boolean is false
The zero value for a string is empty string ==> ""

================================================================================

Type conversion tool:
With this tool, we can change the type of value not type of variable

Ex:
a variable with float type which contains 12.8 (float data type)
we can convert it to integer (12) ==> a variable with int type and a value with int data type

*/

var float_number float64 = 12.8
var int_number int = int(i)

// ---------------------

/*

if in a function you define the parameter type as float, and when you call a function, if you pass it
an integer, in other languages, they change int to float

but in golang, you should only pass the parameter with the type you have defined
or you can change the type of int to float :D

Ex:

func sum(float64 x) {}
sum(float64(10))

*/

// ============================================================================================

// type inference:

func main() {

	var abc int
	j1 := abc

	fmt.Println(j1)
}

/*
when you define a variable which its type is int,
if you define another variable which it gets its primitive data from another
variable, the compiler will set the type of that variable equals the type of the variable
which it gets it's primitive data type from and there is no need to declare the type of the
second variable :D
*/

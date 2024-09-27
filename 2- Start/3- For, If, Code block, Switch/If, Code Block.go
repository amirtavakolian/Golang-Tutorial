package __3

import (
	"fmt"
	"math"
)

/*
# Controlling Statements:

base on different conditions, we have to run different codes
==>
if it was this, do that
if it was that, do this :))

Ex:
if dollar price is below then 100 toman, calculate my products price based on dollar
this is an example of controlling condition.

Ex:
if the student's grade was below then 12, send sms to his/her parents
*/

func sqrt(x int) string {

	result := math.Pow10(x)

	if result > 10 {
		return "great"
	} else {
		return "not bad"
	}
}

// ================================================

func test1() {

	x := math.Sqrt(-1)
	if x == 2 {
		// do something
	}

	// first sqrt the value and put in x
	// then check if its == 2

	// summerized version of above code: no need all above works :D
	if x := math.Sqrt(-1); x == 2 {
		// do something.
	}

	/*
		in the first way, we can have access to x variable outside the if block
		but in second way, we can't access the x variable outside the if block
	*/
}

// ================================================

/*

Code Block:

{  }  ==> this open and close bracket defines the life-time of a variable.
          this means, if you define a variable in the open and close bracket,
          you have access to that variable only inside that open close bracket
          and the opened closed brackets inside that bracket
*/

func test2() {

	y := 5

	if 55 > 10 {

		fmt.Println(y)

		// ==> here i have access to y but no access out of the test2 function

		x := 100
		fmt.Print(x)
	}

	// here, i don't have access to x variable
}

// =========================================

func test3() {

	x := 100

	if x < 50 {
		fmt.Print("lower then 50")
	} else if x == 50 && x <= 70 {
		fmt.Println("between 50 and 70")
	}
}

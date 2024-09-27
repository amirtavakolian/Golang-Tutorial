package __3

import "fmt"

/*

# For Loop:

a tool, which you can use it
when you have lots of data & you want to process these data one by one.

Ex: 100 numbers, you want to check even they are below then 10 or not.
	use for loop to iterate on the data - on each data it does a process
*/

func main() {

	sum := 0

	for i := 0; i <= 100; i += 2 {
		sum += 1
	}

	fmt.Println(sum)
}

// i := 0 ==> i want to start from item 0 of a collection of data

// i += 2 ==> you have to define the steps

// i <= 100 ==> the condition of the for loop
// until i is lower or equal then 100, the loop will continue after that it will exit from loop

// =================================================================================================

// another way to define for loop:

func test() {

	j := 0 // ==> start from 0

	for j <= 100 { // ==> the condition of the for loop
		j += 1 // ==> define the steps
		fmt.Println(j)
	}
}

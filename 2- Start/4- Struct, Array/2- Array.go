package ___Struct__Array

import "fmt"

// a tool/collection/container, which holds several items in allocated ram of the software.
// so you can define a variable with array type for holding several items.

func test() {

	var items [6]int // an array of integer type which has 6 capacity ==> it can hold 6 items - data.

	for key, value := range items {
		fmt.Println(key, value)
	}

	// range, iterates on each item of an array.
	// when using range on an array, for each item of that array, it returns two values
	// ==> index of that item & value of that item
	// index starts from 0
	// "items" variable is a Zero Value array :D :P ==> all items/values are 0

	// ===================================================================

	// give value to the array:
	items = [6]int{1, 2, 3, 4, 5, 6}

	// you can't change the size of the array like below:
	// items [5]int{1,2,3,4,5}

	// you can't do this because the size of the array is 6 not 5
	// (the number of items are different, the first one has 6 items the second one has 5 items)

	// =====================================

	// an array with int type and size 6 is not equal with an array with type int and size 5
	// so, we can't assign/copy them to each-other
}

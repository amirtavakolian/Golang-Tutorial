package ___Struct__Array

import "fmt"

/*

Slice tool is like array but its dynamic size.
we can change the size of the slice at run time :D

2 slices with int type are equal, no matter if they have different sizes
(no matter if they have different amount of items)

*/

func test2() {

	// define an slice
	var ac []int

	// ===========================================

	// built-in functions of golang:

	// add item to slice:
	ac = append(ac, 5)

	// ======================================================================

	/*
			Every program we write has a memory to store its data,
			just like a store has a warehouse to keep its goods.

			Consider the software as the store and the memory allocated to it
			as the warehouse the software uses to store its data.

			Imagine a store that has allocated a part of its warehouse for televisions.
			Ex: it has dedicated 20 meters of warehouse space for televisions,
			which can hold 50 TVs in that 20 meters.

			Now, suppose the store’s sales increase, and it wants to bring in 100 TVs instead of 50.
			Therefore, it needs to increase the warehouse capacity for televisions.

			There are two scenarios:
			either the initial 20 meters has enough space, so they allocate an additional 10 meters next to it
			and place the 100 TVs there, or the warehouse is small,
			and there is no extra space next to the initial 20 meters. In this case, they have to buy another warehouse
			and allocate 50 meters for the televisions instead of 20 meters.

			Two important points:

			Moving is costly and time-consuming.
			They need to find a new place, move the existing TVs to the new location, and continue their work.
			The store owner cannot predict whether sales will decrease or increase in three or six months,
			meaning the warehouse capacity for TVs might change from 20 to 30, or 10, or remain at 20.

			Since it is variable and changes based on various factors, they cannot estimate it accurately.
			So, what is the optimal solution for this person to avoid frequent moving costs? The store might say,
			“I had 100 TVs, and now I have 200, so I had to move.” What should they do? Should they get a place that
			allocates 1000 meters for TVs? Doing so would be very costly, and it is not certain that they will reach
			the capacity of 1000 TVs in their store.

			It is not wise to suddenly allocate 10 or 20 times the capacity. On the other hand,
			it is also not wise to get a warehouse that can hold 22 TVs if the current one holds 20 TVs.

			The likelihood of it filling up quickly and needing to move again is high.

			The optimal solution is to increase the capacity by 1.5 or 2 times whenever there is growth.
			For example, if they have a 20-meter warehouse, they should get a 40-meter warehouse.
			When the 40-meter warehouse is full or nearly full, they should get a 70-80 meter warehouse.

			The memory of the software works exactly the same way. When our software runs and becomes a process,
			the operating system allocates a portion of its 8 or 16 GB RAM to our software, a limited amount.

			Our software must manage this space. If it has 2 GB of RAM, it cannot allocate 1 GB to a specific slice.
			It might not be used, or it might run out of space for other things.

			The best solution is to allocate 5 MB initially, knowing it is empty but will gradually fill up.
			When the 5 MB is nearly full, it is better to increase the space to 10 MB.

			Memory cells are exactly like a warehouse.
			Either there is space next to the memory cells to continue there, or there isn’t.

		    So, I have to move, and moving is very costly ==> if i could stay in the same place, it would be much better for me.

	*/

	/*

		Instead of allocating 5 MB of capacity to the slice, I allocate 10 MB to it but only assign 5 MB.
		When the 5 MB capacity is filled, we increase its capacity to 10 MB.

		The difference is that I already have space there and don’t need to pay the cost of copying that 5 MB
		to another memory address with a larger space, like 10 or 15 MB.

		This is the issue with length and capacity.

		The capacity we actually need becomes the length, and you also have a future prediction
		that it might go up to 10.

		It might even go up to 100 or even 1000,
		but it’s not optimal to give it a capacity of 100 or 1000 from the start.

		It might not fill that 100 or 1000, and on the other hand,
		the software might need RAM for other parts, but we have allocated a capacity of
		1000 for the slice.

		So, you should consider a reasonable limit for slice memory that uses length and capacity.

		You can also not give capacity and just give length.
		if you don't pass capacity and only length, the number for length will be used for capacity :D

	*/

	// the capacity and length, are used for making optimism
	// as we can give initial value to variables, we can define initial size for our slices.
	// Ex: size of slice my_slice must be 5 for now, later i will increase it if needed :D :P

	mySlice := make([]int, 5, 10) // len ==> 5 | cap ==> 10
	fmt.Println(mySlice)

	// length ==> the number of items which currently it has ==> in above code, all the items are zero value ==> search for more info
	// capacity ==>

	d := []int{1, 2, 3, 4, 5}

	// get the capacity of a slice:
	fmt.Println(cap(d))
	fmt.Println(len(d))

	d = append(d, 6) // in above code the length is 5, after append 6, compiler makes the length 5*2=10
	fmt.Println(cap(d))
	fmt.Println(len(d))

	// =============================

	// use slice index to access to its indexes - items:
	fmt.Println(d[0])

	cc := d[1:4] // ==> it just copies index 1, index 2, index 3 ==> index 4 will not copy
	fmt.Println(cc)

	cc1 := d[1:] // from index 1 till the end
	fmt.Println(cc1)

	cc2 := d[:4] // only index 0 1 2 3
	fmt.Println(cc2)

}

package __3

import "fmt"

// there are lots of if, else if in below
// the read ability of the code is so bad

func test4() {

	x := 100

	if x < 50 {
		fmt.Print("lower then 50")
	} else if x == 50 && x <= 70 {
		fmt.Println("between 50 and 70")
	} else if x == 55 && x <= 75 {
		fmt.Println("between 50 and 70")
	} else if x == 60 && x <= 80 {
		fmt.Println("between 50 and 70")
	} else if x == 65 && x <= 85 {
		fmt.Println("between 50 and 70")
	}

	// when the number of else-if goes up, use switch tool
	// below way of switch is used when we want to check the conditions :D :P

	switch {
	case x > 100:
		fmt.Println("test 1")
		break // ==> it will exit of the case if the condition was true.

	case x > 200:
		fmt.Println("test 2")
		fallthrough // it will not exit if the case condition was true and checks the below conditions.

	case x > 200 && x < 100:
		fmt.Println("test 3")
		break

	default:
		fmt.Println("Nothing :D ")
	}

	// ============================================

	// if you want to check the equality of values, use this way of switch

	switch x {
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	default:
		fmt.Println("default :D")
	}
}

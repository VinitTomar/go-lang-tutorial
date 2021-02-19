package flowcontrol

import (
	"fmt"
	"runtime"
	"time"
)

//  ShowSwitch ...
func ShowSwitch() {
	println("======================== Simple switch case 1 =====================")

	fmt.Print("Go is running on ")

	os := runtime.GOOS

	switch os {
	case "darwin":
		println("OS X.")

	case "linux":
		println("Linux.")

	default:
		fmt.Printf("%s. \n", os)
	}

	println("======================== Simple switch case 2 =====================")

	fmt.Print("When is Saturday? ")

	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		println("It is today.")

	case today + 1:
		println("It is tomorrow.")

	case today + 2:
		println("It is in two days.")

	default:
		println("It is too far away")
	}

	println("======================== Simple switch case 3 =====================")

	now := time.Now()

	switch {
	case now.Hour() < 12:
		println("Good morning.")

	case now.Hour() < 17:
		println("Good afternoon.")

	default:
		fmt.Println("Good evening.")
	}
}

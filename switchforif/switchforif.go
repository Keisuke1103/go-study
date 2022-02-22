package switchforif

import (
	"fmt"
)

func Switchforiffunc() {
	thenum := 12
	fmt.Println(thenum)
	for n := 0; n < thenum; n++ {
		switch {
		case n == 0:
		case n == 1:
		case n >= 2:

			if thenum%n == 0 {
				num := thenum / n
				fmt.Printf("を%dで割ると、%d\n", n, num)
			} else {
				num := thenum % n
				fmt.Printf("を%dで割ると、余は%d\n", n, num)
			}
		default:
		}
	}
}

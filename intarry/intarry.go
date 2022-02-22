package main

import (
	"fmt"
)

func main() {
	intvals := [5]int{98, 200, 100, 400, 55}
	// %d は整数を表す %t はブール値を表す
	fmt.Printf("2番目の要素は%d\n", intvals[1])
	fmt.Printf("1番目の要素は%d\n", intvals[0])
	fmt.Printf("4番目の要素は%d\n", intvals[3])
	fmt.Printf("4番目と1番目を足すと%d\n", intvals[3]+intvals[0])
	fmt.Printf("4番目と1番目は同じか%t\n", intvals[3] == intvals[0])
}

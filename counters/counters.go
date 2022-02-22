package counters

import "fmt"

func Counter() func() {
	ctr := 0
	fmt.Println("カウンタの初期化")
	return func() {
		ctr++
		fmt.Println(ctr)
	}
}

func Freecounter(start int) func(int) int {
	ctr := start
	fmt.Printf("フリーカウンタを%dから始めます\n", ctr)
	return func(add int) int {
		fmt.Printf("%dを足して", add)
		ctr += add
		return ctr
	}
}

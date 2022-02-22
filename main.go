package main

import (
	"fmt"
	"main/copying"
	"main/counters"
	"main/switchforif"
	"main/test"
)

func main() {
	fmt.Println(test.Message) // hello world
	switchforif.Switchforiffunc()
	fmt.Println(test.Word) // world
	var dum dummy
	fmt.Println(dum)
	myprint("hoge")
	section4()
	section4on1()
	section6()
}

type dummy struct{}

func myprint(hoge string) (huga string, pword string) {
	huga = " and huga"
	pword = hoge + huga
	return
}

func section4() {
	fmt.Println("-------- Section 4 --------") // world
	countfnc := counters.Counter()
	countfnc()
	countfnc()
	countfnc()
	countfnc2 := counters.Counter()
	countfnc2()
	countfnc()
	fmt.Println("-------- end --------") // world
}

func section4on1() {
	fmt.Println("-------- Section 4-1 --------") // world
	countfnc := counters.Freecounter(10)
	fmt.Println(countfnc(5))
	fmt.Println(countfnc(10))
	fmt.Println(countfnc(21))
	countfnc2 := counters.Freecounter(1)
	fmt.Println(countfnc2(100))
	fmt.Println("-------- end --------") // world
}

func section6() {
	fmt.Println("-------- Section 6 --------") // world
	//copying.Copy()
	copying.Pointout()
	fmt.Println("-------- end --------") // world
}

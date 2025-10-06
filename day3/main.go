package main

import (
	"fmt"
)

const pi float64 = 3.14

const (
	_  = iota
	c1 = 1 << (iota)
	c2
	c3
	c4
	c5
	c6
)

func main() {
	//Using iota
	/*
	fmt.Printf("%d \t %b \n", c1, c1)
	fmt.Printf("%d \t %b \n", c2, c2)
	fmt.Printf("%d \t %b \n", c3, c3)
	fmt.Printf("%d \t %b \n", c4, c4)
	fmt.Printf("%d \t %b \n", c5, c5)
	fmt.Printf("%d \t %b \n", c6, c6)

	fmt.Println("***************")

	fmt.Printf("%d \t %b \n", 1, 1)
	fmt.Printf("%d \t %b \n", 1<<1, 1<<1)
	fmt.Printf("%d \t %b \n", 1<<2, 1<<2)
	fmt.Printf("%d \t %b \n", 1<<3, 1<<3)
	fmt.Printf("%d \t %b \n", 1<<4, 1<<4)
	fmt.Printf("%d \t %b \n", 1<<5, 1<<5)
	fmt.Printf("%d \t %b \n", 1<<6, 1<<6)
	*/
	fmt.Println("Calls another function")
	conversion()
	value_exe5:=exe5()
	fmt.Printf("Value returned by exersise 5 function is: %d\n",value_exe5)
	exec6()
}

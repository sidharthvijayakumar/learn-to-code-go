package main

import (
	"fmt"
	"reflect"
)

func binary(){
	adams := 42
	fmt.Printf("42 in binary is:%#b\n",adams)
	fmt.Printf("42 in hexadecimal is:%#x\n",adams)
	var num1 float32= 4.0
	fmt.Printf("Entered number is of type: %T\n",reflect.TypeOf(num1))
	fmt.Printf("But now its changed to: %s\n",reflect.TypeOf(int(num1)))

}
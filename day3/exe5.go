package main

import (
	"fmt"
)

var global_variale=40

func exe5()(int){
	var num1 int64
	fmt.Printf("Zero value variable is: %d\n",num1)
	num1=42
	fmt.Printf("Short declariation is done for num: %d\n",num1)
	num2,num3,_:=43,50,5.0
	fmt.Printf("Multiple initilisation was done for %d and %d\n",num2,num3)
	fmt.Printf("Var is used to declare variable outside functions:%d\n",global_variale)

	fmt.Println("Black hole was used other values are:",num2,num3)
	return global_variale

}
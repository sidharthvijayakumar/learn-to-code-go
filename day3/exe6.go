package main

import (
	"fmt"
)

func exec6(){
	a,b,c:="Hello",42,42.42
	fmt.Printf("The value of %v and this is of type %T\n",a,a)
	fmt.Printf("The value of %v and this is of type %T\n",b,b)
	fmt.Printf("The value of %v and this is of type %T\n",c,c)

	d,e,f:=747,911,90210

	fmt.Printf("%d in decimal is %d, in binary is %b and in hexadeciam is %#x\n",d,d,d,d)
	fmt.Printf("%d in decimal is %d, in binary is %b and in hexadeciam is %#x\n",e,e,e,e)
	fmt.Printf("%d in decimal is %d, in binary is %b and in hexadeciam is %#x\n",f,f,f,f)

	var num1 uint8=255
	var num2 int8=127
	fmt.Printf("value is nums are:%v is of type %T & %v is of type %T",num1,num1,num2,num2)
}

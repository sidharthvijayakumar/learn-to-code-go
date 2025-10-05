package main

import "fmt"

func main(){
	a :=43
	fmt.Println("Value of variable A is:",a)

	/*In case you do not need a variable use blank identifier
	  Here 59 will be ignored in case if this is not needed*/
	a, b, c, _, d:= 4,10,40,59,100
	println(a,b,c,d)

	/* In case you do not declare a value it will be assigned to 0
	   int: 0
	   string: 
	   bool: false
	   pointer: <nil>
	
	*/
	var(
		num int
		ptr *int 
	) 
	println("The value of number is:",num,ptr)
	num = 43
	ptr=new(int)
	*ptr=40
	println(" The new value of number is:",num,ptr)

}
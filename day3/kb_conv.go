package main

import (
	"fmt"
)

type ByteSize int64

const (
    _ 	= iota // ignore first value by assigning to blank identifier
    KB ByteSize	= 1 << (10 * iota)
    MB
    GB
    TB
)

func conversion (){
	fmt.Printf("KB is %v in decimal and %b in binary \n",KB,KB)
	fmt.Printf("KB is %v in decimal and %b in binary \n",MB,MB)
	fmt.Printf("KB is %v in decimal and %b in binary \n",GB,GB)
	fmt.Printf("KB is %v in decimal and %b in binary \n",TB,TB)
}
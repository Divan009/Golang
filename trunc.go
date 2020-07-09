package main

import "fmt"

func main(){
	var fnum float64
	fmt.Printf("Enter a floating number")
	fmt.Scanln(&fnum)
	fmt.Printf("The number is: %d\n", int(fnum))
}
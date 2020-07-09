package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)


func main(){
	fmt.Printf("Enter the string")
	read := bufio.NewReader(os.Stdin)
	str, _ := read.ReadString('\n')
	
	str = strings.ToLower(str)
	str = strings.TrimSpace(str)
	
	if strings.HasPrefix(str, "i") && 
	strings.Contains(str, "a") &&
	strings.HasSuffix(str, "n"){
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}

}
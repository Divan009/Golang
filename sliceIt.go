package main

import (
	"fmt"
	"sort"
	"os"
	"strings"
	"strconv"
)

func main(){
	//an empty slice of size 3
	num := ""
	sli := make([]int,0, 3)

	for true{
		fmt.Println("Enter a num")
		fmt.Scan(&num)

		if strings.ToUpper(num) == "X"{
			fmt.Println("Exit")
			os.Exit(0)
		}
		x, err := strconv.Atoi(num)
		if err != nil {	
			continue
		}
		sli = append(sli, x)
		
		sort.Ints (sli)
		fmt.Println(sli)
	}
}
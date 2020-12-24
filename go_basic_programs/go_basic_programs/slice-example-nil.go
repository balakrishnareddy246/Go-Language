package main 

import "fmt"

func main() {

	var numbers [] int

	printSlice(numbers)

	if(numbers == nil) {
		fmt.Printf("slice is nil Balakrishna reddy")
	}

}

func printSlice(x [] int) {

	fmt.Println("len = %d cap = %d slice = %v\n", len(x), cap (x),x)
}
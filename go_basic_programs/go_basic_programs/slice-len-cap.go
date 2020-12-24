package main 

import "fmt"

func main() {

	balakrishna := []int{2,3,5,7,11,13}
	printSlice(balakrishna)
	
	balakrishna = balakrishna[:0]
	printSlice(balakrishna)

	balakrishna = balakrishna[1:2]
	printSlice(balakrishna)
	balakrishna = balakrishna[4:]

	printSlice(balakrishna)

}

func printSlice(balakrishna []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(balakrishna), cap(balakrishna), balakrishna)

}
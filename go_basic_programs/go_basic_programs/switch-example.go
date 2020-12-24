package main 

import "fmt"

func main(){
 
i :=3

switch i {
   case 1:
	    fmt.Println("one")
    case 2:
	    fmt.Println("two")
    case 3:
	    fmt.Println("three")
}
j := 5
switch j {
case 1:
	fmt.Println("one")
case 2:
	fmt.Println("two")
case 3:
	fmt.Println("three")
default:
	fmt.Println("Not one, two, or three")
}
switch {
	case j == 5: 
	fmt.Println(j, "is equal to 5")
case j < 5:
	fmt.Println(j, "is less than 5")
case j > 5:
	fmt.Println(j, "is greater than 5")
}
}
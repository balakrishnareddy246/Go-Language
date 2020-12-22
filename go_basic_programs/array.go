package main 

import "fmt"

func main () {

	var int_arr [5]int
	fmt.Println(int_arr)

	var bool_arr [10]bool

	fmt.Println(bool_arr)

	var str_arr [20]string
    fmt.Println(str_arr)
	int_arr[0] = 5

	int_arr[1] = 10
	int_arr[3] = 30
	int_arr[4] = 20
  fmt.Println(int_arr)

  fmt.Println(int_arr[0])
  a :=[5]int{1,2,3,4,5}
  fmt.Println(len(a))
  fmt.Println(len(bool_arr))

  var aa[5][5]int
  count := 0 
  for i := 0; i < 5; i++ {
	  for j:=0; j< 5; j++ {
		  aa[i][j] = count
		  count = count + 1
	  }

	}
	fmt.Println(aa)
}


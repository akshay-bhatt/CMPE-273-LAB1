package main

import "fmt"


func main() {

fmt.Print("Enter a number for the fibonacci series: ")
var my_no int
fmt.Scan( &my_no)
var temp_no int
  for i :=1 ; i < my_no ; i ++ {
  temp_no = fibonacci_prog(i)
  }
 fmt.Println(temp_no)
}

func fibonacci_prog(pass_no int ) int{
 if pass_no <= 1 {
	return pass_no
 }
 
 return fibonacci_prog(pass_no - 1) + fibonacci_prog(pass_no - 2)
 
}

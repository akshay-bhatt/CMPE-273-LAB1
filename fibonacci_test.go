
package main

import "testing"

func Test_fibonacci(t *testing.T) {
var test_no int
 test_no = fibonacci_prog(4)
 if test_no != 2{
 t.Error("Expected 2 , got " , test_no)
 }
 
}

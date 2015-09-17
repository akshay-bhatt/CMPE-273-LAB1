package main

import (
    "fmt"
    "time"
  )

func main() {
c1 := make(chan string)
c2 := make(chan string)

go func() {
for {
c1 <- "Make Candy"
time.Sleep(time.Second * 1)
}
}()

go func() {
for {
c2 <- "Pack Candy "
time.Sleep(time.Second * 1)
}
}()

go func() {
for {
select {
case msg1 := <- c1:
fmt.Println(msg1)
fmt.Println(time.Now())
case msg2 := <- c2:
fmt.Println(msg2)
fmt.Println(time.Now())
case  <- time.After(time.Second):
fmt.Println("timeOut")
}
}
}()
var input string
fmt.Scanln(&input)
}
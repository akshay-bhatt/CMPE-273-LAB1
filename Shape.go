
package main

import "fmt"
import "math"


type Main_Shape interface {
    Area() float64
    Perimeter() float64
}


type shape_rect struct {
    rect_width, rect_height float64
}
type shape_circle struct {
    circ_radius float64
}


func (r_obj shape_rect) Area() float64 {
fmt.Print("Area of the rectangle is: ")
    return r_obj.rect_width * r_obj.rect_height
}
func (r_obj shape_rect) Perimeter () float64 {
fmt.Print("Perimeter of the rectangle is: ") 
 return 2*r_obj.rect_width + 2*r_obj.rect_height
}


func (c_obj shape_circle) Area() float64 {
 fmt.Print("Area of the circle is: ")
    return math.Pi * c_obj.circ_radius * c_obj.circ_radius
}
func (c_obj shape_circle) Perimeter() float64 {
 fmt.Print("Perimeter of the circle is: ")   
   return 2 * math.Pi * c_obj.circ_radius
}

func measure(g_obj Main_Shape) {
   // fmt.Println(g)

  
    fmt.Println(g_obj.Area())
    fmt.Println(g_obj.Perimeter())
}

func main() {
 fmt.Print("Enter the height and the width of the retangle: ")
var h float64
var w float64
fmt.Scan( &h)
fmt.Scan( &w)
r := shape_rect{rect_width: w, rect_height: h}
measure(r)


fmt.Print("Enter the radius of the circle: ")
var cr float64
fmt.Scan( &cr)
c:= shape_circle{circ_radius: cr}
measure(c)
}

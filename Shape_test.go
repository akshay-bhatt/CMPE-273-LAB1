
package main

import "testing"



func Test_Shape(t *testing.T ) {

    r := shape_rect{rect_width: 20, rect_height: 4}
    c := shape_circle{circ_radius: 10}


    measure(r)
    measure(c)
    
}

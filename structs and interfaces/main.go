package main

import (
	"fmt"
)



type Shape interface{
	Area() float64
	Perimeter() float64
}
type Rectangle struct{
	Width float64
	Height float64
}

func (r Rectangle) Area()float64{
	return r.Width * r.Height
}
func (r Rectangle) Perimeter()float64{
	return 2 * (r.Width + r.Height)
}

var s Shape = Rectangle{Width:10,Height:5}

func main(){
	fmt.Println("Area :",s.Area())
	fmt.Println("Perimeter :",s.Perimeter())
}

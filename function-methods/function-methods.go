package main

import (
   "fmt" 
   "math" 
)

/* define a circle */
type Circle struct {
   x,y,radius float64
}

/* define a method for circle */
func(c Circle) area() float64 {
   return math.Pi * c.radius * c.radius
}

func main(){
   c := Circle{x:0, y:0, radius:5}
   fmt.Printf("Circle area: %f\n", c.area())
}

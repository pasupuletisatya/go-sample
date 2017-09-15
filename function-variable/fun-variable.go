package main

import (
   "fmt" 
   "math" 
)

func main(){
   /* declare a function variable */
   getSquareRoot := func(x float64) float64 {
      return math.Sqrt(x)
   }

   /* use the function */
   var k float64 = getSquareRoot(20)
   fmt.Println(k)
   fmt.Println(getSquareRoot(9))
}

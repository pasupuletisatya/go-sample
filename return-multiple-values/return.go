package main

import "fmt"

func main() {
   /* local variable definition */
   var a int = 100
   var b int = 200
//   var ret1, ret2 int

   /* calling a function to get max value */
   //ret1, ret2 = max(a, b)
   r1, r2 := max(a, b)

   //fmt.Printf( "Max value is : %d and %d\n", ret1, ret2 )
   fmt.Printf( "Max value is : %d and %d\n", r1, r2 )
}

/* function returning the max between two numbers */
func max(num1, num2 int) (int,int) {
   /* local variable declaration */
   var result int

   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result , 100
}

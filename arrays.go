package main
import "fmt"

func main() {
   var a [5]int;
   fmt.Println("array declared: ", a )

   a[4] = 100
   fmt.Println("value at index 4: ", a[4])
   fmt.Println("array set: ", a )

   fmt.Println("len of array: ", len(a))

   b := [5]int {1,2,3,4,5}
   fmt.Println("array initialized: ", b )

   var c [2][3]int;
   for i:=0;i<2;i++ {
       for j:=0;j<3;j++ {
           c[i][j] = j
       }
   }
   fmt.Println("2D array", c)
}

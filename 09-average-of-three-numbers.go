//write a function that calculates the average of three floating point numbers and returns the average and the sum.
package main
import "fmt"

func averageAndSum(x, y, z float64) (float64, float64){
  sum := x + y + z
  avg := sum / 3
  return avg, sum
}

func main(){
  avg, sum := averageAndSum(1.0, 2.0, 3.0)
  fmt.Println("Average:", avg, "Sum:", sum)
}

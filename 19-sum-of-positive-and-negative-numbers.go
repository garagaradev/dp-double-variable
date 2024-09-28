//write a function that sums the positive and negative numbers separately from a slice
//and return both sums.
package main
import "fmt"

func sumPosNeg(numbers []int) (int, int){
  positiveSum, negativeSum := 0,0
  for _, num := range numbers {
    if num > 0 {
      positiveSum += num
    }
    if num < 0 {
      negativeSum += num
    }
  }
  return positiveSum, negativeSum
}

func main(){
  posSum, negSum := sumPosNeg([]int{1,-2,3,-4,5})
  fmt.Println("Positive sum:",posSum,"Negative sum:",negSum)
}

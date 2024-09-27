//write a function that finds the minimum and maximum values in a slice of integers
package main
import "fmt"

func minMaxSlice(numbers []int) (int, int){
  min, max := numbers[0], numbers[0]
  for _, number := range numbers {
    if number < min {
      min = number
    }
    if number > max {
      max = number
    }
  }
  return min, max
}

func main(){
  numbers := []int{20, 20, 30, 45, 10, 5, 90, 1000, 500}
  fmt.Println("numbers:", numbers)
  min, max := minMaxSlice(numbers)
  fmt.Println("min:",min," & max:", max)
}

// Write a function that returns the first and last elements of a slice of integers
package main
import "fmt"

func firstAndLast(numbers []int) (int, int){
  if len(numbers) == 0 {
    return 0, 0
  }

  return numbers[0], numbers[len(numbers)-1]
}

func main(){
  numbers := []int{1,2,3,4,5,6,7}
  fmt.Println("numbers:",numbers)
  
  first, last := firstAndLast(numbers)
  fmt.Println("First element:", first, ". Last element:", last)
}

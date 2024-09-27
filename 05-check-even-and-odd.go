// write a function that checks if a number is even or odd and returns both results
package main
import "fmt"

func checkEvenOdd(n int) (bool, bool){
  return n%2 == 0, n%2 != 0
}

func main(){
  num := 5
  fmt.Println("num is ", num)
  isEven, isOdd := checkEvenOdd(num)
  fmt.Println("Is even? ", isEven, ". Is odd? ", isOdd)
}

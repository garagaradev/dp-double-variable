//write a function that checks whether a number is divisible by both 3 and 5
package main
import "fmt"

func checkDivisibility(n int)(bool, bool){
    divisibleBy3 := n % 3 == 0
    divisibleBy5 := n % 5 == 0
    return divisibleBy3, divisibleBy5
}

func main(){
  n := 15
  fmt.Println("Number:",n)
  div3, div5 := checkDivisibility(n)
  fmt.Println(n, "is divisible by 3?", div3,".",n," is divisible by 5?", div5)
  
}

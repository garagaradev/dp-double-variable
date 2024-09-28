// Write a function that generates a random even and odd numbers
package main
import (
  "fmt"
  "math/rand"
  "time"
)

func randomEvenOdd()(int, int){
  rand.Seed(time.Now().UnixNano())
  even := rand.Intn(50)*2
  odd  := rand.Intn(50)*2+1
  return even, odd
}

func main(){
  even, odd := randomEvenOdd()
  fmt.Println("Even:",even,"Odd:",odd)
}

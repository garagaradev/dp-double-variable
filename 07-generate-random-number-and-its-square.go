// Write a function that generates a random integer and returns the number and its square
package main
import (
  "fmt"
  "math/rand"
  "time"
)

func randomNumberAndSquare() (int, int){
  rand.Seed(time.Now().UnixNano()) // Seed the random number generator
  num := rand.Intn(100) // Random number between 0 and 99
  return num, num * num
}

func main(){
  number, square := randomNumberAndSquare()
  fmt.Println("number:", number, ". Its square:", square)
}

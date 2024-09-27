//write a function that reverses a string and returns the reversed string along with its length.
package main
import "fmt"

func reverseString(s string) (string, int) {
    reversed := ""
    for _, char := range s {
    reversed = string(char) + reversed
  }
  return reversed, len(reversed)
}

func main(){
  reversed, length := reverseString("hello")
  fmt.Println("Reversed string:", reversed, ". Length:", length)
}

// Write a function that counts the number of characters in a string and returns the count along with the string itself.
package main
import "fmt"

func countCharacters(s string) (int, string){
  return len(s), s
}

func main(){
  count, original := countCharacters("Hello, World!")
  fmt.Println("Character count:",count,". Original string:", original)
}

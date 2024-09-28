// write a function that splits a string into its first and last word
package main
import (
  "fmt"
  "strings"
)

func firstAndLastWord(s string) (string, string){
  words := strings.Fields(s)
  if len(words) == 0 {
    return "",""
  }
  return words[0],words[len(words)-1]
}

func main(){
  words := "lorem ipsum dolor sit amet consectetur adipscing elite"
  first, last := firstAndLastWord(words)
  fmt.Println("The words:", words)
  fmt.Println("First word:",first,". Last word:",last)
}


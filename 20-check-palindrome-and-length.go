//write a function that checks if a string is a palindrome and returns
//whether it is, along with its length
package main
import "fmt"

func isPalindrome(s string)(bool, int){
  length := len(s)
  for i:=0;i<length/2;i++{
    if s[i] != s[length-i-1]{
      return false, length
    }
  }
  return true, length
}

func main(){
  word := "risetovotesir"
  palindrome, length := isPalindrome(word)
  fmt.Printf("QUESTION: Is '%s', a palindrome? \nANSWER:'%t'\nQUESTION: How much is the length?\nANSWER:'%d'",word,palindrome,length)
}

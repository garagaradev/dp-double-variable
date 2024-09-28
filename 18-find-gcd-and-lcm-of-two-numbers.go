//write a function that finds the greatest common divisor (GCD) 
//and least common multiple (LCM) of two numbers.
package main
import "fmt"

func gcd(a,b int)int {
  for b!= 0 {
    a, b = b, a%b
  }
  return a
}

func lcm(a, b int)int{
  return a * b / gcd (a,b)
}

func gcdAndLcm(a,b int)(int,int){
  return gcd(a,b), lcm(a,b)
}

func main(){
  var num1,num2 int
  fmt.Print("Input the number:")
  fmt.Scan(&num1)
  fmt.Print("Input the number:")
  fmt.Scan(&num2)
  gcd, lcm := gcdAndLcm(num1,num2)
  fmt.Println("GCD:",gcd,".LCM:",lcm)
}

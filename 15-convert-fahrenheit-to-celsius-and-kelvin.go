// Write a function that converts Fahrenheit to Celsius and Kelvin
package main
import "fmt"

func fahrenheitToCelsiusAndKelvin(f float64)(float64, float64){
  celsius := (f-32) * 5/9
  kelvin := celsius + 273.15
  return celsius, kelvin
}

func main(){
  var fahrenheit float64 = 98.6
  fmt.Println("fahrenheit:", fahrenheit)
  celsius, kelvin := fahrenheitToCelsiusAndKelvin(fahrenheit)
  fmt.Println("Celsius:",celsius,". Kelvin:",kelvin)
}

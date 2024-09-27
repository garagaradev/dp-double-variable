//write a function that converts kilometers to miles and returns both values.
package main
import "fmt"

func kilometersToMiles(kilometers float64) (float64, float64){
  miles := kilometers * 0.621371
  return kilometers, miles
}

func main(){
  var num float64 = 10
  fmt.Println("km:", num)
  km, mi := kilometersToMiles(num)
  fmt.Println("Kilometers:", km, ". Miles:", mi)
}

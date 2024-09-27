// write a function that returns the circumference and area of a circle given its radius.
package main
import (
  "fmt"
  "math"
)

func circleProperties(radius float64) (float64, float64){
  circumference := 2 * math.Pi * radius
  area := math.Pi * radius * radius
  return circumference, area
}

func main(){
  var radius float64 = 5
  fmt.Println("radius:", radius)

  circumference, area := circleProperties(radius)
  fmt.Println("if the radius is ", radius, ", the circumference is:", circumference, " and the area is:", area)
}

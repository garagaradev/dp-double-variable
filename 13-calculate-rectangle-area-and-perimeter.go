//write a function that calculates the area and perimeter of a rectangle given its length and width
package main
import "fmt"

func rectangleProperties(length, width float64) (float64, float64){
  area := length * width
  perimeter := 2 * (length + width)
  return area, perimeter
}

func main(){
  var length, width float64 = 10, 5
  fmt.Println("length:",length,", width:",width)
  area, perimeter := rectangleProperties(length, width)
  fmt.Println("Area:",area,". Perimeter:",perimeter)
}

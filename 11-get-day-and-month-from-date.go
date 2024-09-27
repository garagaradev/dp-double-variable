//write a function that returns the day and the month of a given date 
package main
import (
  "fmt"
  "time"
)

func getDayAndMonth(date time.Time) (int, int){
  return date.Day(), int(date.Month())
}

func main(){
  day, month := getDayAndMonth(time.Now())
  fmt.Println("Day:", day, "Month:", month)
}

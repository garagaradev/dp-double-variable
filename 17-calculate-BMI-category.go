package main

import "fmt"

func calculateBMI(weight, height float64) (float64, string){
    bmi := weight / (height * height)
    var category string
  switch {
  case bmi <  18.5:
    category = "Underweight"
  case bmi < 24.9:
    category = "Normal Weight"
  case bmi < 29.9:
    category = "Overweight"
  default:
    category = "Obese"
  }
  return bmi, category
}

func main(){
  bmi, category := calculateBMI(56.0,170.0)
  fmt.Println("BMI:",bmi,"Category:",category)
}



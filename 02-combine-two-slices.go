// write a function that combines two slices of integers and returns the combined slice
// and its length.
package main
import "fmt"

func combineSlices(s1, s2 []int) ([]int, int){
  combined := append(s1, s2...)
  return combined, len(combined)
}

func main(){
  slice1 := []int{1,2,3}
  slice2 := []int{4,5,6}
  fmt.Println("slice1:",slice1)
  fmt.Println("slice2:",slice2)
  combinedSlice, length := combineSlices(slice1, slice2)
  fmt.Printf("Combined Slices:%d, slices length:%d\n",combinedSlice, length)
}

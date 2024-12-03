package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error){
  if e!= nil {
    panic(e)
  }
}

func check_left_right(arr []int) {
  fmt.Println(arr[0],arr[1],arr[2],arr[3],arr[4],arr[5],arr[6],arr[7],arr[8])
}

func day1(xd []string){
  var right []int
  var left []int
  for index, element := range xd{
    if index % 2 == 0{
      left_tmp, _ := strconv.Atoi(element)
      left = append(left, left_tmp)
    }
    if index % 2 == 1{
      right_tmp, _ := strconv.Atoi(element)
      right = append(right, right_tmp)
    }
  }
  slices.Sort(left)
  slices.Sort(right)
  diff := 0
  for i:= 0; i< len(left) && i< len(right) ; i++{
    if left[i] > right[i]{
      diff += left[i]- right[i]
    }else{
      diff += right[i]- left[i]
    }
  }
  fmt.Println("part1",diff)
  similarty := 0
  for Index, element := range left{
    occurance := 0
    for rIndex, rElement := range right{
      if element == rElement{
        occurance +=1
      }else if rIndex > Index && element != rElement{
        break
      }
    }
    similarty += occurance * element
  }
  fmt.Println("part2", similarty)
}

func main(){
  dat, err := os.ReadFile("input_01.txt")
  check(err)
  xd := strings.Split(string(dat),"\n")
  //fmt.Println(xd)
  var input []string
  for i:= 0; i< len(xd); i++{ 
    f := func(c rune) bool {
		  return c == ' '
	  }
    d := strings.FieldsFunc(xd[i], f)
    input = append(input, d...)
  }
  day1(input)
}


package main

import (
	"fmt"
	"os"
	"strings"
  "strconv"
)

func check(e error){
  if e!= nil {
    panic(e)
  }
}
func part1(xd string) bool{
  lineStr := strings.Split(xd, " ")
  var line []int
  for _, element := range lineStr{
    line_tmp, _ := strconv.Atoi(element)
    line = append(line, line_tmp)
  }
  fmt.Println(line)
  isSafe := true
  var isDescending bool
  for i:= 1; i < len(line) ; i++{
    diff := line[i-1] - line[i]
    if i == 1{
      if line[i-1]<line[i]{
        isDescending = false
      }else{
        isDescending = true
      }
    }
    if diff == 0{
      isSafe = false
      break
    }
    if isDescending{
      if line[i-1]< line[i]{
        isSafe = false
        break
      }else{
        if diff !=2 && diff != 1 {
          isSafe = false
          break
        }
      }
    }else{
      if line[i-1] > line[i]{
        isSafe = false
        break
      }else{
        if diff != -2 && diff != -1{
          isSafe = false
          break
        }
      }
    }
  }
  fmt.Println(isSafe)
  return isSafe
}

func main(){
  dat, err := os.ReadFile("input_02.txt")
  check(err)
  xd := strings.Split(string(dat),"\n")
  //fmt.Println(xd)
  safe :=0
  fmt.Println(xd[len(xd)-1])
  for _, element := range xd{
    isSafe := part1(element)
    if isSafe{
      safe +=1
    }
  }
  fmt.Println(safe)
}

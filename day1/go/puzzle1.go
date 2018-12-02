package main

import (
    "fmt"
    "strings"
    "strconv"
    "io/ioutil"
)

func main() {

  dat, err := ioutil.ReadFile("../input1.txt")
  if err != nil {
    panic(err)
  }

  var data_str = strings.TrimSpace(string(dat));
  var lines = strings.Split(data_str, "\n")

  sum := 0
  for i := 0; i < len(lines); i++ {
    int_val, _ := strconv.Atoi(lines[i])
    sum += int_val
	}
  fmt.Println("Sum:", sum)

}

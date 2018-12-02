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

  frequencies_len := len(lines)
  frequencies := make([]int, frequencies_len)
  for i := 0; i < frequencies_len; i++ {
    int_val, _ := strconv.Atoi(lines[i])
    frequencies[i] = int_val;
	}

  index := 0
  sum := 0
  sum_set := make(map[int]bool)
  for true {
    sum += frequencies[index]
    _, found := sum_set[sum]
    sum_set[sum] = true
    if found {
      break
    }

    index += 1
    if index == frequencies_len {
      index = 0
    }
  }
  fmt.Println("Sum:", sum)
}

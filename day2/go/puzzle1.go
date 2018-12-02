package main

import (
    "fmt"
    "strings"
    "io/ioutil"
)

func getLines() []string {
  dat, err := ioutil.ReadFile("../input1.txt")
  if err != nil {
    panic(err)
  }

  var data_str = strings.TrimSpace(string(dat));
  var lines = strings.Split(data_str, "\n")
  return lines
}

func countLetters(containerId string) (bool, bool) {
  twoCount := false
  threeCount := false
  countMap := make(map[rune]int)
  for _, ch :=range containerId {
    chCount, found := countMap[ch]
    if !found {
      countMap[ch] = 1
    } else {
      countMap[ch] = (chCount + 1)
    }
  }
  for _, count := range countMap {
    if count == 2 {
      twoCount = true
    }
    if count == 3 {
      threeCount = true
    }
    if twoCount && threeCount {
      break
    }

  }
  return twoCount, threeCount
}

func countDiff(line1 string, line2 string) (bool, int) {
  count := 0
  diffIndex := -1
  for i := 0; i < len(line1); i++ {
    if line1[i] != line2[i] {
      count++
      diffIndex = i
      if count > 1 {
        return false, -1
      }
    }
  }
  return count == 1, diffIndex
}

func part1(lines []string) int {
  twoCount := 0
  threeCount := 0
  for _, line := range lines {
    two, three := countLetters(line)
    if two {
      twoCount++
    }
    if three {
      threeCount++
    }
  }
  checkSum := twoCount * threeCount
  return checkSum
}

func part2(lines []string) string {
  for i, line1 := range lines {
    for j:=i+1; j < len(lines); j++ {
      line2 := lines[j]
      found, diffIndex := countDiff(line1, line2)
      if found {
        return line1[:diffIndex] + line1[diffIndex+1:]
      }
    }
  }
  return ""
}

func main() {
  lines := getLines()

  fmt.Printf("------- Part 1 -------\n")
  checkSum := part1(lines)
  fmt.Printf("%d\n", checkSum);

  fmt.Printf("------- Part 2 -------\n")
  fmt.Printf("%s \n", part2(lines))

}

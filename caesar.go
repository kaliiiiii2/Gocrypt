package main

import (
  "fmt"
  "strings"
)

func isUpper (s string) bool {
  return s == strings.ToUpper(s)
}

func isLower (s string) bool {
  return s == strings.ToLower(s)
}

func main () {
  var plainText string = "Hello World";
  var key int = 3;

  for _, char := range plainText {
    var shifted int = int(char) + key

    if string(char) == " " {
      fmt.Printf(" ")

    } else if isUpper(string(char)) {
      if shifted > int('Z') {
        shifted -= 26
      }
    } else if isLower(string(char)) {
      if shifted > int('z') {
        shifted -= 26
      }
    }
    fmt.Printf("%s", string(shifted));
  }
}

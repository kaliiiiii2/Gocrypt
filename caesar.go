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
  var result string;

  for _, char := range plainText {
    var shifted int = int(char) + key

    if string(char) == " " {
      result += " "

    } else if isUpper(string(char)) {
      if shifted > int('Z') {
        shifted -= 26
      }
      result += string(shifted);

    } else if isLower(string(char)) {
      if shifted > int('z') {
        shifted -= 26
      }
      result += string(shifted);
    }
  }
  
  fmt.Printf ("Encryption: %s\n", result);

  var cipher string = result;
  var result_2 string;
  
  for _, char := range cipher {
    var shifted int = int(char) - key

    if string(char) == " " {
      result_2 += " "

    } else if isUpper(string(char)) {
      if shifted < int('A') {
        shifted += 26
      }
      result_2 += string(shifted);

    } else if isLower(string(char)) {
      if shifted < int('a') {
        shifted += 26
      }
      result_2 += string(shifted);
    }
  }
  
  fmt.Printf ("Decryption: %s", result_2);
}

package main

import (
  "fmt"
  "strings"
)

var (
  rot13 string = "abcdefghijklmnopqrstuvwxyz";
  rot18 string = "abcdefghijklmnopqrstuvwxyz0123456789"
)

func isUpper (s string) bool {
  return s == strings.ToUpper(s); 
}

func isLower (s string) bool {
  return s == strings.ToLower(s); 
}

func encrypt (plaintText string, key string) string{
  var result string;

  for _, char := range plaintText { 
    if string(char) == " " {
      result += " "  

    } else if key == "rot13"{
      var is_upper bool = isUpper(string(char));
      var char  string = strings.ToLower(string(char));

      
      if strings.Contains(rot13, char) {
        var index int = strings.Index (rot13, char)
        if is_upper {
          result += strings.ToUpper(string(rot13[(index-13 + 26) % 26])) 
        } else {
          result += string(rot13[(index-13 + 26) % 26])
        }
      } 
    } else if key == "rot18" {
  		if strings.Contains(rot18, string(char)) {
				index := strings.Index(rot18, string(char))
				result += string(rot18[(index+18)%36])
			} else {
				result += string(char)
			}
    }
  }
  return result
}

func main () {

  var result string = encrypt("Hw336 W693v", "rot18"); 
  fmt.Printf("\n Hello World %s", result)
}

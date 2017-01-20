package main

import (
  "fmt"
)

func main() {
  //Initialize letters slice
  letters := []string{"A","B","C","D","E",
    "F","G","H","I","J","K","L","M",
    "N","O","P","Q","R","S","T",
    "U","V","W","X","Y","Z"}

  //Create the vigenere square
  var square [26][26]string

  k := 0
  //loop through rows
  for i := 0; i < 26; i++ {
    //loop through columns
    for j := 0; j < 26; j++ {
      k = (i + j) % 26
      square[i][j] = letters[k]
    }
  }

  fmt.Println("Square:")
  fmt.Println(square)
}

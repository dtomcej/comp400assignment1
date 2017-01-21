package main

import (
  "bufio"
  "fmt"
  "math/rand"
  "os"
  "strings"
  "time"
)

func main() {
  //Initialize letters slice and map of letters
  letters := []string{"A","B","C","D","E",
    "F","G","H","I","J","K","L","M",
    "N","O","P","Q","R","S","T",
    "U","V","W","X","Y","Z"}
  letters_map := make(map[string]int)
  for i := 0; i < len(letters); i++ {
    letters_map[letters[i]] = i
  }

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

  fmt.Println("Do you want to:")
  fmt.Println("1: Encrypt Message")
  fmt.Println("2: Decrypt Message")

  reader := bufio.NewReader(os.Stdin)
  char, _, err := reader.ReadRune()

  if err != nil {
  fmt.Println(err)
  }

  switch char {
  case '1':
    key := ""
    //initiate a new seed on runtime, and use it as seed for random generator
    source := rand.NewSource(time.Now().UnixNano())
    gen := rand.New(source)
    fmt.Println("Random Keys:")
    //print 4 random keys of length 8
    for i := 0; i < 4; i++ {
      key = ""  //reset key
      for j := 0; j < 8; j++ {
        key += letters[gen.Intn(26)]
      }
    //Print out the key
    fmt.Printf("Key %v: %v\n", i, key)
    }

    fmt.Println("Please enter a key to use for encryption:")

    reader := bufio.NewReader(os.Stdin)
    raw_key, _ := reader.ReadString('\n')
     // print out confirmation
    fmt.Printf("You have entered key: %v" ,raw_key)

    fmt.Println("Please enter a plaintext message:")

    raw_plaintext, _ := reader.ReadString('\n')
     // print out confirmation
    fmt.Printf("You have entered text: %v" ,raw_plaintext)

    clean_key := strings.TrimSpace(strings.ToUpper(strings.Replace(raw_key, " ", "", -1)))
    clean_plaintext := strings.TrimSpace(strings.ToUpper(strings.Replace(raw_plaintext, " ", "", -1)))

    //if key is longer than plaintext, truncate
    if len(clean_key) > len(clean_plaintext) {
      clean_key = clean_key[:len(clean_plaintext)]
    }

    key_text := ""
    for len(key_text) < len(clean_plaintext) {
      key_text += clean_key
    }
    if len(key_text) > len(clean_plaintext) {
      key_text = key_text[:len(clean_plaintext)]
    }

    ciphertext := ""
    for i := 0; i < len(clean_plaintext); i++ {
      ciphertext += square[letters_map[string(key_text[i])]][letters_map[string(clean_plaintext[i])]]
    }
    fmt.Printf("Cleaned and formatted Key: %v\n", clean_key)
    fmt.Printf("Cleaned and formatted Plaintext: %v\n", clean_plaintext)
    fmt.Printf("Cleaned and formatted Keytext: %v\n", key_text)
    fmt.Printf("Cleaned and formatted Ciphertext: %v\n", ciphertext)
    break
  case '2':
    fmt.Println("Please enter a key to use for decryption:")

    reader := bufio.NewReader(os.Stdin)
    raw_key, _ := reader.ReadString('\n')
     // print out confirmation
    fmt.Printf("You have entered key: %v" ,raw_key)

    fmt.Println("Please enter a ciphertext message:")

    raw_ciphertext, _ := reader.ReadString('\n')
     // print out confirmation
    fmt.Printf("You have entered text: %v" ,raw_ciphertext)

    clean_key := strings.TrimSpace(strings.ToUpper(strings.Replace(raw_key, " ", "", -1)))
    clean_ciphertext := strings.TrimSpace(strings.ToUpper(strings.Replace(raw_ciphertext, " ", "", -1)))

    //if key is longer than ciphertext, truncate
    if len(clean_key) > len(clean_ciphertext) {
      clean_key = clean_key[:len(clean_ciphertext)]
    }

    key_text := ""
    for len(key_text) < len(clean_ciphertext) {
      key_text += clean_key
    }
    if len(key_text) > len(clean_ciphertext) {
      key_text = key_text[:len(clean_ciphertext)]
    }

    plaintext := ""
    for i := 0; i < len(clean_ciphertext); i++ {
      for j := 0; j < 26; j++ {
        if square[letters_map[string(key_text[i])]][j] == string(clean_ciphertext[i]) {
          plaintext += letters[j]
          break
        }
      }
    }
    fmt.Printf("Cleaned and formatted Key: %v\n", clean_key)
    fmt.Printf("Cleaned and formatted Ciphertext: %v\n", clean_ciphertext)
    fmt.Printf("Cleaned and formatted Keytext: %v\n", key_text)
    fmt.Printf("Cleaned and formatted Plaintext: %v\n", plaintext)
    break
  }
}

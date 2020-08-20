package binaryNcoder

import (
  "fmt"
  "errors"
)

func Encode(data string) string {
  var buffer string
  arr := []byte(data)
  for _, b := range arr {
    for _, q := range fmt.Sprintf("%08b", b) {
      buffer += fmt.Sprintf("%s", string(q))
    }
  }
  return buffer
}

func bitStringToBytes(s string) ([]byte, error) {
    b := make([]byte, (len(s)+(8-1))/8)
    for i := 0; i < len(s); i++ {
        c := s[i]
        if c < '0' || c > '1' {
            return nil, errors.New("value out of range")
        }
        b[i>>3] |= (c - '0') << uint(7-i&7)
    }
    return b, nil
}

func Decode(data string) string {
  b, err := bitStringToBytes(data)
  if err != nil { fmt.Println(data, err) }
  return string(b)
}

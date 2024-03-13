package md5

import (
  MD5 "crypto/md5"
  "fmt"
)

func Encode(text string) string{
  var result [16]byte = MD5.Sum([]byte(text))
  return fmt.Sprintf("%x",result)
}
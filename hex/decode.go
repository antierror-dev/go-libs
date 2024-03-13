package hex

import(
  HEX "encoding/hex"
)

func Decode(text string) string{
  result ,_ := HEX.DecodeString(text)
  return string(result)
}
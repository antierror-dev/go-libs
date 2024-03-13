package hex

import(
  HEX "encoding/hex"
)

func Encode(text string) string{
  return HEX.EncodeToString([]byte(text))
}
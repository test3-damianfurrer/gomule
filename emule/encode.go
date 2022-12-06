package emule

import (
	"fmt"
)

func encodeByteMsg(protocol byte,msgcode byte,body []byte) []byte {
  size := uint32ToByte(uint32(len(body)))
  return []byte{protocol,
    size[0], size[1], size[2], size[3],
    msgcode,
		body...}
}

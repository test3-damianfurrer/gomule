package emule

import (
	"fmt"
)

func encodeByteMsg(protocol byte,msgcode byte,body []byte) []byte {
	bodysize := len(body)
	size := uint32ToByte(uint32(bodysize+1))
	buf := [size+6]byte{protocol, size[0], size[1], size[2], size[3],msgcode}
		//body...
	for i := 0; i < len(size); i++ {
		buf[i+6] = body[i]
	}
	return buf
}

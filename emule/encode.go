package emule

import (
	"fmt"
)

func encodeByteMsg(protocol byte,msgcode byte,body []byte) []byte {
	size := uint32ToByte(uint32(len(body)))
	buf := [size+6]byte{protocol, size[0], size[1], size[2], size[3],msgcode}
		//body...
	for i := 0; i < len(size); i++ {
		data[i+6] = body[i]
	}
	return buf
}

package emule

import (
)

func encodeByteMsg(protocol byte,msgcode byte,body []byte) []byte {
	bodysize := len(body)
	size := uint32ToByte(uint32(bodysize+1))
	buf := make([]byte,bodysize+6)
	buf[0] = protocol
	buf[1] = size[0]
	buf[2] = size[1]
	buf[3] = size[2]
	buf[4] = size[3]
	buf[5] = msgcode
	for i := 0; i < len(bodysize); i++ {
		buf[i+6] = body[i]
	}
	return buf
}

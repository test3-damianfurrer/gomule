package emule

import (
)

func encodeByteMsg(protocol byte,msgcode byte,body []byte) []byte {
	bodysize := len(body)
	sizebytes := uint32ToByte(uint32(bodysize+1))
	buf := make([]byte,bodysize+6)
	buf[0] = protocol
	buf[1] = sizebytes[0]
	buf[2] = sizebytes[1]
	buf[3] = sizebytes[2]
	buf[4] = sizebytes[3]
	buf[5] = msgcode
	for i := 0; i < bodysize; i++ {
		buf[i+6] = body[i]
	}
	return buf
}

func encodeByteString(str string) []byte {
	slen:=len(str)
	buf := make([]byte,slen+2)
	sizebytes := uint16ToByte(uint16(slen))
	buf[0] = sizebytes[0]
	buf[1] = sizebytes[1]
	for i := 0; i < slen; i++ {
		buf[i+2] = str[i]
	}
}

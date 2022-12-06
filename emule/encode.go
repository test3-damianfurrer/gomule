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
	return buf
}


//func encodeByteTag(ttype byte, tagname []byte, tagvalue []byte, specialdesignator byte){
//	
//}

func encodeByteTagString(tagname []byte, tagvalue string) []byte {
	return encodeByteTag(2,tagname,stringToByte(tagvalue))
}

func encodeByteTagInt(tagname []byte, tagvalue uint32) []byte {
	return encodeByteTag(3,tagname,uint32ToByte(tagvalue))
}

/*func encodeByteTagInt(tagname []byte, tagvalue float) []byte {
	return encodeByteTag(4,tagname,uint32ToByte(tagvalue))
}*/

func encodeByteTag(ttype byte, tagname []byte, tagvalue []byte) []byte {
	//buflen=len(tagname)+len(tagvalue)+1
	buf := make([]byte,0)
	buf = append(buf, ttype)
	buf = append(buf, tagname...)
	buf = append(buf, tagvalue...)
	return buf
	
	/*
	Login example: 
	vers tag:   [3 1 0 17 60 0 0 0]
	port tag:   [3 1 0 32 29 3 0 0] -> 0x20 (port should be 0x0f, 15)
	flag tag:   [3 1 0 251 128 13 4 3] //seems to be some other tag (should be 0x20)
	name tag: [2 1 0 1 (strlen)(string) ]
	2/3 = type string/int
	[1 0] = bytes for tag name
	17(0x11) = tag name value for Version Tag

	Offer files example(somehow reversed?):
	[1 4 0 116 101 115 116] //simple
	[0 0 1 4 0 116 101 115 116 2 5 0 73 109 97 103 101 1 0 3]
	initial 0 0 - not sure, marking complex search and maybe something else
	type, value(len, str), tagname
	string type, 1 0 3 = Filtype tag name
	value = "Image"(5)
	
	*/
}

func encodeByteTagNameInt(val byte) []byte {
	buf := make([]byte,1)
	buf[0]=val
	return encodeByteTagName(buf)
}
func stringToByte(val string) []byte {
	strlen:=len(val)
	buf := make([]byte,strlen)
	for i := 0; i < strlen; i++ {
		buf[i] = val[i]
	}
	return buf
}
func encodeByteTagNameStr(val string) []byte {
	return encodeByteTagName(stringToByte(val))
}

func encodeByteTagName(nbuf []byte) []byte {
	blen:=len(nbuf)
	buf := make([]byte,blen+2)
	sizebytes := uint16ToByte(uint16(blen))
	buf[0] = sizebytes[0]
	buf[1] = sizebytes[1]
	for i := 0; i < blen; i++ {
		buf[i+2] = nbuf[i]
	}
	return buf
}

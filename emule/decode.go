import (
	"fmt"
)

type OneTag struct {
  Type byte
  NameByte byte
  NameString string
  Value []byte
  ValueLen uint16
}

func readString(pos int, buf []byte)(bread int, ret string) {
  readname:=0
  bread=2
  bread += int(byteToInt16(buf[pos:pos+2]))
  ret := fmt.Sprintf("%s",buf[pos+2:bread])
}

func readTag(pos int, buf []byte)(bread int, ret *OneTag) {
  ret := OneTag{Type: buf[pos], NameString: ""}
  //byteToUint16
  bread=3
  namelen := byteToInt16(buf[pos+1:pos+bread])
  
  if namelen == uint16(1) {
    ret.NameByte = buf[pos+3]
    readname = 1
  } else {
    readname, ret.NameString = readString(pos+3,buf)
  }
  bread+=readname
  
  //[3 1 0 17 60 0 0 0]
  
  switch ret.Type {
    case byte(2): //varstring
      ret.ValueLen = uint16(byteToInt16(buf[pos+bread:pos+bread+2]))
      bread += 2
      ret.Value = buf[pos+bread:pos+bread+ret.ValueLen]
      bread+=ret.ValueLen
    case byte(3): //uint32
      ret.ValueLen = 4
      buf[pos+bread:pos+bread+4]
      bread += 4
    case byte(4): //float
      ret.ValueLen = 4 //implemented?
    default:
      fmt.Println("Error decoding Tag, unknown tag datatype!",ret.Type)
    }
  
  
}

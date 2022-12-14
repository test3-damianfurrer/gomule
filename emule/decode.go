package emule

import (
	"fmt"
)
//TODO: check buf len on all and prevent read > len(buf)
type OneTag struct {
  Type byte
  NameByte byte
  NameString string
  Value []byte
  ValueLen uint16
}

type constrainttype byte
const (
	C_NONE constrainttype = iota
	C_MAIN
	C_AND
	C_OR
	C_NOT
	C_CODEC
	C_MINSIZE
	C_MAXSIZE
	C_FILETYPE
	C_FILEEXT
)

type Constraint struct {
	Type constrainttype
	Value []byte
	Left *Constraint
	Right *Constraint
}
func enumNumberConstraint(one byte, two byte, three byte, four byte) constrainttype {
	switch one {
		case 3:
			switch two {
				case 1:
					switch three {
						case 0:
							switch four {
								case 211:
									return C_MAXSIZE
								default:
									return C_NONE
							}
						default:
							return C_NONE
					}
				default:
					return C_NONE
			}
		default:
			return C_NONE
	}
}

func enumStringConstraint(one byte, two byte, three byte) constrainttype {
	switch one {
		case 1:
			switch two {
				case 0:
					switch three {
						case 213:
							return C_CODEC
						case 4:
							return C_FILEEXT
						case 3:
							return C_FILETYPE
						default:
							return C_NONE
					}
				default:
					return C_NONE
			}
		default:
			return C_NONE
	}
	return C_NONE
}

func readConstraints(pos int, buf []byte)(readb int,ret *Constraint){
	readb=pos
	fmt.Println("Const read pos:", pos)
	fmt.Println("Const read buf:", buf)
	//var main Constraint
	switch buf[readb] {
		case 0x0:
			readb+=1
			switch buf[readb] {
				case 0x0:
					ret = &Constraint{Type: C_AND}
					fmt.Println("Debug AND identifier")
				/* 2 bytes, ignore so far [ 0x01 0x00 ] [ 0x02 0x00 ] .. how to differenciate from [1/2] 0 1 [mutiple of 10 byte string]
				case 0x100:
[					main = Constraint{Type: C_OR}
				case 0x200:
					main = Constraint{Type: C_NOT}
					*/
				default:
					fmt.Println("ERROR expected either AND/OR/NOT identifier")
					readb-=pos
					return
			}
			readb+=1
			readsub, subret := readConstraints(readb,buf)
			readb+=readsub
			ret.Left = subret
			readsub, subret = readConstraints(readb,buf)
			ret.Right = subret
			readb+=readsub
		case 0x1:
			readb+=1
			strlen:=int(ByteToUint16(buf[readb:readb+2]))
			readb+=2
			ret = &Constraint{Type: C_MAIN, Value: buf[readb:readb+strlen]}
			readb+=strlen
			fmt.Println("Debug Main Constraint")
		case 0x2: //string value
			readb+=1
			strlen:=int(ByteToUint16(buf[readb:readb+2]))
			readb+=2
			ret = &Constraint{Value: buf[readb:readb+strlen]}
			readb+=strlen
			ret.Type = enumStringConstraint(buf[readb],buf[readb+1],buf[readb+2])
			if ret.Type == C_NONE {
				fmt.Println("ERRROR unrecognized string constraint type!",buf[readb:readb+3])
			}
			readb+=3
			
		case 0x3: //int value
			readb+=1
			ret = &Constraint{Value: buf[readb:readb+4]}
			readb+=4
			ret.Type = enumNumberConstraint(buf[readb],buf[readb+1],buf[readb+2],buf[readb+3])
			if ret.Type == C_NONE {
				fmt.Println("ERRROR unrecognized number constraint type!",buf[readb:readb+4])
			}
			readb+=4
		default:
			fmt.Printf("ERROR: unexpected byte: 0x%x \n",buf[readb])
			readb+=1
	}
	readb-=pos
	return

}

func ReadTags(pos int, buf []byte, tags int,debug bool)(totalread int, ret []*OneTag){
	index := pos
	totalread = 0
	fmt.Println("TAGS BUF:",buf[pos:pos+50])
	for i := 0; i < tags; i++ {
		bread, tag := ReadTag(index,buf,debug)
		totalread += bread
		index += bread
		ret = append(ret,tag)
	}
	return
}

func readString(pos int, buf []byte)(bread int, ret string) {
  fmt.Println("readstring!",buf[pos-3:len(buf)])
  bread=2
  bread += int(ByteToUint16(buf[pos:pos+2]))
  ret = fmt.Sprintf("%s",buf[pos+2:bread])
  return
}

func ReadTag(pos int, buf []byte, debug bool)(bread int, ret *OneTag) {
  dpos = pos + 50
  if dpos > len(buf){
	  dpos = len(buf)
  }
	
  fmt.Println("TAG BUF:",buf[pos:dpos])
  if debug {
    fmt.Println("readtag! at ",pos)
  }
  ret = &OneTag{Type: buf[pos], NameString: ""}
  bread=3
  readname:=0
  namelen := ByteToUint16(buf[pos+1:pos+bread])
  if debug {
    fmt.Println("name tag len",namelen)
  }
  
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
      ret.ValueLen = ByteToUint16(buf[pos+bread:pos+bread+2])
      bread += 2
      ret.Value = buf[pos+bread:pos+bread+int(ret.ValueLen)]
      bread+=int(ret.ValueLen)
    case byte(3): //uint32
      ret.ValueLen = 4
      ret.Value = buf[pos+bread:pos+bread+4]
      bread += 4
    case byte(4): //float
      ret.ValueLen = 4
      ret.Value = buf[pos+bread:pos+bread+4]
      bread += 4
    default:
      fmt.Println("Error decoding Tag, unknown tag datatype!",ret.Type)
    }
  
  return
}

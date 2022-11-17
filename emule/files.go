package emule

import (
	"fmt"
	"net"
	libdeflate "github.com/4kills/go-libdeflate/v2"
)

//type Mode int
// The constants that specify a certain mode of compression/decompression
//const (
//	ModeDEFLATE Mode = iota
//	ModeZlib
//	ModeGzip
//)
func prcofferfiles(buf []byte, conn net.Conn, debug bool, blen int) {
	//30 bytes more: [2 1 0 1 15 0 66 111 100 121 98 117 105 108 100 101 114 46 109 112 52 3 1 0 2 104 11 112 0 2]
	// =
	// [2 1 0 1] len[15 0 ] Bodybuilder.mp4 [3 1 0 2 104 11 112 0 2]
	
	//30 bytes more: [2 1 0 1 50 0 116 104 101 46 115 105 109 112 115 111 110 115 46 115 48 50 101 49 48 46 105 110 116 101]
	// [2 1 0 1] len 50 
  count := byteToInt32(buf[0:4]) //cant be more than 200 by spec
  if debug {
    fmt.Println("DEBUG: prcofferfiles")
    fmt.Println("DEBUG: files:", count)
  }
  iteration := 1
  byteoffset := uint32(4)
  for{
    if byteoffset >= uint32(blen) {
	    break;
    }
    if debug {
      fmt.Println("DEBUG: byteoffset", byteoffset)
      fmt.Println("DEBUG: iteration", iteration)
    }
    fuuid := fmt.Sprintf("%x-%x-%x-%x-%x-%x-%x-%x",
		buf[byteoffset+0:byteoffset+2], buf[byteoffset+2:byteoffset+4], 
		buf[byteoffset+4:byteoffset+6], buf[byteoffset+6:byteoffset+8],
		buf[byteoffset+8:byteoffset+10], buf[byteoffset+10:byteoffset+12], 
		buf[byteoffset+12:byteoffset+14], buf[byteoffset+14:byteoffset+16])
    fmt.Println("DEBUG: 1.  filehash:", fuuid)
    fmt.Println("DEBUG: 1. client id:", buf[byteoffset+16:byteoffset+20])
    fmt.Println("DEBUG: 1. client port:", buf[byteoffset+20:byteoffset+22])
    itag := byteToInt32(buf[byteoffset+22:byteoffset+26])
    fmt.Println("DEBUG: 1. tag count:", itag)
	  //skip 4 [2 1 0 1] 
    strlen := uint32(byteToInt16(buf[byteoffset+30:byteoffset+32]))
    strbuf := buf[byteoffset+32:byteoffset+32+strlen]
    str := fmt.Sprintf("%s",strbuf)
    	  
    fmt.Println("DEBUG: 1. File name:", str)
    //[3 1 0 2]
    fsize := byteToUint32(buf[byteoffset+32+strlen+4:byteoffset+32+strlen+8])
    fmt.Println("DEBUG: 1. File size:", fsize)
    //[2 1 0 3]
    strlentype := uint32(byteToInt16(buf[byteoffset+32+strlen+12:byteoffset+32+strlen+14]))
    strbuf = buf[byteoffset+32+strlen+14:byteoffset+32+strlen+14+strlentype]
    str = fmt.Sprintf("%s",strbuf)
    fmt.Println("DEBUG: 1. File type:", str)
    byteoffset = byteoffset+32+strlen+14+strlentype
    //fmt.Println("DEBUG: 30 bytes more:", buf[byteoffset+36+strlen+14+strlentype:byteoffset+36+strlen+14+strlentype+30])
    iteration+=1
  }
}
func offerfiles(buf []byte, protocol byte, conn net.Conn, debug bool, n int) {
  if debug {
	fmt.Println("DEBUG: Client offers Files / Keep alive")
	fmt.Printf("DEBUG: File offering protocol 0x%02x\n", protocol)
  }
  bufcomp := buf[1:n]
  if protocol == 0xd4 {
	var blen int = 0
 	var decompressed []byte
	dc, err := libdeflate.NewDecompressor() //not recomended to create a new instance each, but also not possible to use the same simultaniously
  	if err != nil {
		fmt.Println("ERROR libdeflate:", err.Error())
		return
  	}
  	fmt.Println("DEBUG: decompressing")
  	blen, decompressed, err = dc.Decompress(bufcomp, nil, 1)
	dc.Close()
	  fmt.Println("DEBUG: after decompressing")
	if err != nil {
		fmt.Println("ERROR decompress:", err.Error())
		fmt.Println("ERROR: uncompressed len", blen)
	  	fmt.Println("ERROR: uncompressed buf 10", decompressed[0:10])
		return
	}
  	fmt.Println("DEBUG: uncompressed len", blen) //9927 vs 9928-1 compressed? lol??? There might be something off
  	fmt.Println("DEBUG: uncompressed buf 10", decompressed[0:10])
	prcofferfiles(decompressed, conn, debug, blen)
  } else if protocol == 0xe3 {
	prcofferfiles(bufcomp, conn, debug, n-1)
  } else {
	  fmt.Println("Error: offerfiles: worong protocol")
  }

}

func filesources(buf []byte, protocol byte, conn net.Conn, debug bool, n int) {
	//type=buf[0]
  if debug {
    fmt.Println("DEBUG: Client looks for File Sources")
    //fmt.Println("DEBUG: filehash:", buf[1:n])
    //fmt.Println("DEBUG: 16lehash:", buf[1:17])
    //fmt.Println("DEBUG: 16revhas:", buf[n-16:n])
  }
}


func listservers(buf []byte, protocol byte, conn net.Conn, debug bool, n int) {
	//type=buf[0]
  if debug {
    fmt.Println("DEBUG: Get list of servers")
    //fmt.Println("DEBUG: listservers")
  }
}

func searchfiles(buf []byte, protocol byte, conn net.Conn, debug bool, n int) {
	//type=buf[0]
  
  
  
  //if debug {
if 1==2 {
    fmt.Println("DEBUG: Client looks for Files")
    //fmt.Println("DEBUG: searchfiles")
    fmt.Println("DEBUG: buf full query:", buf[1:n])
    if(buf[1] == 0x01) {
	fmt.Println("DEBUG: simple search")
    	strlen := byteToInt16(buf[2:4])
    	fmt.Println("DEBUG: strlen:", strlen)
    	fmt.Println("DEBUG: strlen buf:", buf[2:4])
    	fmt.Println("DEBUG: buf string:", buf[4:4+strlen])
    	strbuf := buf[4:4+strlen]
    	str := fmt.Sprintf("%s",strbuf)
	fmt.Println("DEBUG: str:", str)
    } else {
	fmt.Println("DEBUG: complex search")
	strlen := byteToInt16(buf[4:6])
    	fmt.Println("DEBUG: strlen:", strlen)
    	fmt.Println("DEBUG: strlen buf:", buf[4:6])
	fmt.Println("DEBUG: buf string:", buf[6:6+strlen])
	strbuf := buf[6:6+strlen]
    	str := fmt.Sprintf("%s",strbuf)
	fmt.Println("DEBUG: str:", str)
	    fmt.Println("DEBUG: buf other:", buf[6+strlen:n])
    }
    //fmt.Println("DEBUG: buf query:", buf[1:n])
	  
	  //buf query: [1 5 0 101 109 117 108 101]
	  //emule, len 5
//DEBUG: buf full query: [1 5 0 101 109 117 108 101]
//DEBUG: strlen: 5
//DEBUG: strlen buf: [5 0]
//DEBUG: buf full query: [0 0 1 5 0 101 109 117 108 101 2 3 0 68 111 99 1 0 3]
//emule + type texts
	  
	  //search emule with type texts
	  //[0 0 1 5 0 101 109 117 108 101 2 3 0 68 111 99 1 0 3]
	  
	  //buf other: [2 3 0 68 111 99 1 0 3]
	  //68 111 99 -> what is that? ASCII: "Doc"
	  //search with type text and ending pdf: other:
	  //[0 0 2 3 0 68 111 99 1 0 3 2 3 0 112 100 102 1 0 4]
	  //[0 0 2] -> ?
	  //[3 0 68 111 99] -> Doc, len 3
	  //[1 0 3] -> ? , also in search with type only
	  //[2] -> ?
	  //[3 0 112 100 102] -> pdf, len 3
	  //[1 0 4] -> ?
  }
}
		
func requestcallback(buf []byte, protocol byte, conn net.Conn, debug bool, n int) {
	//type=buf[0]
  if debug {
    fmt.Println("DEBUG: Client looks for another to callback")
  }
}

func udpfilesources(buf []byte, protocol byte, conn net.Conn, debug bool, n int) {
	//type=buf[0]
  if debug {
    fmt.Println("DEBUG: UDP Client looks for File Sources")
  }
}


package emule

import (
	"fmt"
	"net"
	"errors"
	libdeflate "github.com/4kills/go-libdeflate/v2"
)

//type Mode int
// The constants that specify a certain mode of compression/decompression
//const (
//	ModeDEFLATE Mode = iota
//	ModeZlib
//	ModeGzip
//)

func offerfiles(buf []byte, protocol byte, conn net.Conn, debug bool, n int) {
  if debug {
	fmt.Println("DEBUG: Client offers Files / Keep alive")
	fmt.Printf("DEBUG: File offering protocol 0x%02x\n", protocol)
  }
if 1 == 2 {
//initial file offering seems to be always of size 9224 
  var blen int = 0
  var decompressed []byte
  //[]byte decompressed = nil
	//type=buf[0]
//it's compressed ...
  dc, err := libdeflate.NewDecompressor() //not recomended to create a new instance each, but also not possible to use the same simultaniously
  if err != nil {
	fmt.Println("ERROR libdeflate:", err.Error())
	return
  }
  fmt.Println("DEBUG: decompressing")
  //if 1 != 1 {
  //blen, decompressed, err = dc.DecompressZlib(buf[1:n], nil)
	//libdeflate.Mode
  bufcomp := buf[1:n]
  blen, decompressed, err = dc.Decompress(bufcomp, nil, 1)
  fmt.Println("DEBUG: after decompressing")
  if err != nil {
	fmt.Println("ERROR decompress:", err.Error())
	fmt.Println("ERROR: uncompressed len", blen)
  	fmt.Println("ERROR: uncompressed buf 10", decompressed[0:10])
	return
  }
  
  fmt.Println("DEBUG: uncompressed len", blen)
  fmt.Println("DEBUG: uncompressed buf 10", decompressed[0:10])
  //}
  if 1 != 1 {
  count := byteToInt32(buf[1:5]) //spec says, can't be more than 200, but is 4 bytes? The resulting number seems utter garbage
  if debug {
    fmt.Println("DEBUG: type:", buf[0])
    fmt.Println("DEBUG: files:", count)
    //fmt.Println("DEBUG: metadata:", buf[5:n])
    fuuid := fmt.Sprintf("%x-%x-%x-%x-%x-%x-%x-%x",
		buf[5:7], buf[7:9], buf[9:11], buf[11:13], buf[13:15], buf[15:17], buf[17:19],
		buf[19:21])
    fmt.Println("DEBUG: 1.  filehash:", buf[5:21])
    fmt.Println("DEBUG: 1.  filehash:", fuuid)
    fmt.Println("DEBUG: 1. client id:", buf[21:25])
    cport := byteToInt16(buf[25:27])
    fmt.Println("DEBUG: 1. client port:", buf[25:27])
    fmt.Println("DEBUG: 1. client port:", cport)
    itag := byteToInt32(buf[27:31])
    fmt.Println("DEBUG: 1. tag count:", buf[27:31])
    fmt.Println("DEBUG: 1. tag count:", itag)
    fmt.Println("DEBUG: 10 bytes more:", buf[31:41])
  }
  }
  dc.Close()
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


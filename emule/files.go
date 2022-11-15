package emule

import (
	"fmt"
	"net"
)

func offerfiles(buf []byte, protocol byte, conn net.Conn, debug bool, n int) {
	//type=buf[0]
  count := byteToInt32(buf[1:5]) //spec says, can't be more than 200
  if debug {
    fmt.Println("DEBUG: type:", buf[0])
    fmt.Println("DEBUG: files:", count)
    count := byteToInt32(buf[5:9])
    fmt.Println("DEBUG: files(tst):", count)
    fmt.Println("DEBUG: filecnt buf:", buf[1:5])
    //fmt.Println("DEBUG: metadata:", buf[5:n])
    fuuid := fmt.Sprintf("%x-%x-%x-%x-%x-%x-%x-%x",
		buf[5:7], buf[7:9], buf[9:11], buf[11:13], buf[13:15], buf[15:17], buf[17:19],
		buf[19:21])
    fmt.Println("DEBUG: 1. filehash:", buf[5:21])
    fmt.Println("DEBUG: 1. filehash:", fuuid)
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

func filesources(buf []byte, protocol byte, conn net.Conn, debug bool, n int) {
	//type=buf[0]
  if debug {
    fmt.Println("DEBUG: filehash:", buf[1:n])
    fmt.Println("DEBUG: 16lehash:", buf[1:17])
  }
}


func listservers(buf []byte, protocol byte, conn net.Conn, debug bool, n int) {
	//type=buf[0]
  if debug {
    fmt.Println("DEBUG: listservers")
  }
}

func searchfiles(buf []byte, protocol byte, conn net.Conn, debug bool, n int) {
	//type=buf[0]
  strlen := byteToInt32(buf[2:4])
  strbuf := buf[4:4+strlen]
  str := fmt.Sprintf("%s",strbuf)
  if debug {
    fmt.Println("DEBUG: searchfiles")
    fmt.Println("DEBUG: buf query:", buf[1:n])
    fmt.Println("DEBUG: buf string:", buf[4:4+strlen])
    fmt.Println("DEBUG: strlen:", strlen)
    fmt.Println("DEBUG: strlen buf:", buf[2:4])
    fmt.Println("DEBUG: str:", str)
    //fmt.Println("DEBUG: buf query:", buf[1:n])
	  
	  //buf query: [1 5 0 101 109 117 108 101]
	  //emule, len 5
  }
}
		
func requestcallback(buf []byte, protocol byte, conn net.Conn, debug bool, n int) {
	//type=buf[0]
  if debug {
    fmt.Println("DEBUG: requestcallback")
  }
}

func udpfilesources(buf []byte, protocol byte, conn net.Conn, debug bool, n int) {
	//type=buf[0]
  if debug {
    fmt.Println("DEBUG: listservers")
  }
}


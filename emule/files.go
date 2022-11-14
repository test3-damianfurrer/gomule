package emule

import (
	"fmt"
	"net"
)

func offerfiles(buf []byte, protocol byte, conn net.Conn, debug bool, n int) {
	//type=buf[0]
  count := byteToInt32(buf[1:5])
  if debug {
    fmt.Println("DEBUG: type:", buf[0])
    fmt.Println("DEBUG: files:", count)
    fmt.Println("DEBUG: filecnt buf:", buf[1:5])
    fmt.Println("DEBUG: metadata:", buf[5:n])
  }
}

func filesources(buf []byte, protocol byte, conn net.Conn, debug bool, n int) {
	//type=buf[0]
  if debug {
    fmt.Println("DEBUG: filehash:", buf[1:n])
    fmt.Println("DEBUG: 16lehash:", buf[1:17])
  }
}

package emule

import (
	"fmt"
	"net"
)

func offerfiles(buf []byte, protocol byte, conn net.Conn, debug bool, n int) {
	//type=buf[0]
  count := byteToInt32(buf[1:5])
  if debug {
    fmt.Println("DEBUG: files:", count)
    fmt.Println("DEBUG: metadata:", buf[5:n])
  }
}

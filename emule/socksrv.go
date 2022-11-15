/*
 * Copyright (C) 2013 Deepin, Inc.
 *               2013 Leslie Zhai <zhaixiang@linuxdeepin.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package emule

import (
	"fmt"
	"io"
	"net"

	sam "github.com/eyedeekay/sam3/helper"
)

type SockSrv struct {
	Host     string
	Port     int
	Debug    bool
	I2P      bool
	SAM      string
	SAMPort  int
	listener net.Listener
}

func NewSockSrv(host string, port int, debug bool) *SockSrv {
	return &SockSrv{
		Host:  host,
		Port:  port,
		Debug: debug}
}

func (this *SockSrv) read(conn net.Conn) (buf []byte, protocol byte, err error, buflen int) {
	protocol = 0xE3
	buf = make([]byte, 5)
	err = nil
	var n int = 0

	n, err = conn.Read(buf)
	if err != nil {
		if err != io.EOF {
			fmt.Println("ERROR:", err.Error())
		}
		return
	}
	if buf[0] == 0xE3 {
		protocol = 0xE3
	}
	if this.Debug {
		fmt.Printf("DEBUG: protocol 0x%02x\n", protocol)
	}
	size := byteToUint32(buf[1:n])
	if this.Debug {
		fmt.Printf("DEBUG: size %v -> %d\n", buf[1:n], size)
	}
	buf = make([]byte, size)
	n, err = conn.Read(buf)
	buflen = n
	return
}

func (this *SockSrv) respConn(conn net.Conn) {
	if this.Debug {
		fmt.Printf("DEBUG: %v connected\n", conn.RemoteAddr())
	}
	for {
		buf, protocol, err, buflen := this.read(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("DEBUG: %v disconnected\n", conn.RemoteAddr())
			}
			return
		}
		if this.Debug {
			fmt.Printf("DEBUG: type 0x%02x\n", buf[0])
		}
		if buf[0] == 0x01 {
			if this.Debug {
				fmt.Println("DEBUG: Login")
			}
			login(buf, protocol, conn, this.Debug)
		} else if buf[0] == 0x14 {
			if this.Debug {
				fmt.Println("DEBUG: Get list of servers")
			}
			listservers(buf, protocol, conn, this.Debug, buflen)
		} else if buf[0] == 0x15 {
			if this.Debug {
				fmt.Println("DEBUG: Client offers Files / Keep alive")
			}
			offerfiles(buf, protocol, conn, this.Debug, buflen)
		} else if buf[0] == 0x16 {
			if this.Debug {
				fmt.Println("DEBUG: Client looks for Files")
			}
			searchfiles(buf, protocol, conn, this.Debug, buflen)
		} else if buf[0] == 0x19 {
			if this.Debug {
				fmt.Println("DEBUG: Client looks for File Sources")
			}
			filesources(buf, protocol, conn, this.Debug, buflen)
		} else if buf[0] == 0x1c {
			if this.Debug {
				fmt.Println("DEBUG: Client looks for another to callback")
			}
			requestcallback(buf, protocol, conn, this.Debug, buflen)
		} else if buf[0] == 0x9a {
			if this.Debug {
				fmt.Println("DEBUG: UDP Client looks for File Sources")
			}
			udpfilesources(buf, protocol, conn, this.Debug, buflen)
		}
	}
}

func (this *SockSrv) yoursam() string {
	return fmt.Sprintf("%s:%d", this.SAM, this.SAMPort)
}

func (this *SockSrv) Start() {
	if this.I2P {
		ln, err := sam.I2PListener("go-imule-servr", this.yoursam(), "go-imule-server")
		if err != nil {
			fmt.Println("ERROR:", err.Error())
			return
		}
		this.listener = ln
		fmt.Printf("Starting server %s:%d\n", this.Host, this.Port)

		for {
			conn, err := this.listener.Accept()
			if err != nil {
				fmt.Println("ERROR:", err.Error())
				continue
			}
			go this.respConn(conn)
		}
	} else {
		ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Host, this.Port))
		if err != nil {
			fmt.Println("ERROR:", err.Error())
			return
		}
		this.listener = ln
		fmt.Printf("Starting server %s:%d\n", this.Host, this.Port)

		for {
			conn, err := this.listener.Accept()
			if err != nil {
				fmt.Println("ERROR:", err.Error())
				continue
			}
			go this.respConn(conn)
		}
	}
}

func (this *SockSrv) Stop() {
	defer this.listener.Close()
	return
}

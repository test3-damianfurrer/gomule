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
	"errors"

	sam "github.com/eyedeekay/sam3/helper"
	"database/sql"
	libdeflate "github.com/4kills/go-libdeflate/v2"
	_ "github.com/go-sql-driver/mysql"
)

type SockSrv struct {
	Host     string
	Port     int
	Debug    bool
	Ssname   string
	Ssdesc	 string
	Ssmsg	 string
	I2P      bool
	SAM      string
	SAMPort  int
	SupportGzip		bool
	SupportNewTags		bool
	SupportUnicode		bool
	SupportRelSearch	bool
	SupportTTagInteger	bool
	SupportLargeFiles	bool
	SupportObfuscation	bool
	SQL      bool
	SqlDriver string
	SqlUser  string
	SqlPW    string
	SqlAddr	 string
	SqlPort	 int
	SqlDB    string
	db       *sql.DB
	listener net.Listener
}
type SockSrvClient struct {
	Conn net.Conn
	Comp	   libdeflate.Compressor
	DeComp	   libdeflate.Decompressor
}

func (this *SockSrv) getTCPFlags() (ret uint32) {
	ret = uint32(0)
	if this.SupportGzip{
		ret +=  uint32(0x00000001)
	}
	if this.SupportNewTags{
		ret += uint32(0x00000008)
	}
	if this.SupportUnicode{
		ret += uint32(0x00000010)
	}
	if this.SupportRelSearch{
		ret += uint32(0x00000040)
	}
	if this.SupportTTagInteger{
		ret += uint32(0x00000080)
	}
	if this.SupportLargeFiles{
		ret += uint32(0x00000100)
	}
	if this.SupportObfuscation{
		ret += uint32(0x00000400)
	}
	if this.Debug {
		fmt.Printf("DEBUG: used Serverflags: %b\n",ret)
	}
	return
}

func NewSockSrv(host string, port int, debug bool) *SockSrv {
	return &SockSrv{
		Host:  host,
		Port:  port,
		Debug: debug}
}

func (this *SockSrv) read(conn net.Conn) (buf []byte, protocol byte, err error, buflen int) {
	//possible protocols:
	//0xe3 - ed2k
	//0xc5 - emule
	//0xd4 -zlib compressed
	protocol = 0xE3
	buf = make([]byte, 5)
	err = nil
	var n int = 0

	n, err = conn.Read(buf)
	if err != nil {
		/*if err != io.EOF {
			fmt.Println("ERROR:", err.Error())
			}
		*/
		return
	}
	if buf[0] == 0xE3 {
		protocol = 0xE3
	} else if buf[0] == 0xD4 {
		protocol = 0xD4
	} else if buf[0] == 0xC5 {
		protocol = 0xC5
	} else {
		fmt.Printf("ERROR: unsuported protocol 0x%02x\n", protocol)
		err = errors.New("unsuported protocol")
		return
	}
	if this.Debug {
		fmt.Printf("DEBUG: selected protocol 0x%02x(by byte 0x%02x)\n", protocol, buf[0])
	}
	size := ByteToUint32(buf[1:n])
	//if this.Debug {
	//	fmt.Printf("DEBUG: size %v -> %d\n", buf[1:n], size)
	//}
	buf = make([]byte, 0)
	toread := size
	var tmpbuf []byte
	for{
		if toread > 1024  {
			tmpbuf = make([]byte, 1024)
		} else {
			tmpbuf = make([]byte, toread)
		}
		n, err = conn.Read(tmpbuf)
		if err != nil {
			fmt.Println("ERROR: on read to buf", err.Error())
			//return
		}
		buf = append(buf, tmpbuf[0:n]...)
		if n < 0 {
			fmt.Println("WARNING: n (conn.Read) < 0, some problem")
			n = 0
		}
		toread -= uint32(n)
		if toread <= 0 {
			break;
		}
	}
	//buf = make([]byte, size)
	//n, err = conn.Read(buf)
	//if err != nil {
	//	fmt.Println("ERROR: on read to buf", err.Error())
	//	//return
	//}
	n = int(size-toread)
	if this.Debug {
		fmt.Printf("DEBUG: size %d, n %d\n", size, n)
	}
	buflen = n
	return
}

func (this *SockSrv) respConn(conn net.Conn) {
	//var chigh_id uint32
	//var cport int16
	
	//test
	var err error
	uhash := make([]byte, 16)
	client := SockSrvClient{Conn: conn}
	
	
	client.DeComp, err = libdeflate.NewDecompressor()
	if err != nil {
		fmt.Println("ERROR libdeflate Decompressor:", err.Error())
		return
	}
	client.Comp, err = libdeflate.NewCompressor()
	if err != nil {
		fmt.Println("ERROR libdeflate Compressor:", err.Error())
		return
	}
	
	if this.Debug {
		fmt.Printf("DEBUG: %v connected\n", conn.RemoteAddr())
	}
	for {
		buf, protocol, err, buflen := this.read(conn)
		if err != nil {
			if err == io.EOF {
				if this.Debug {
				    fmt.Printf("DEBUG: %v disconnected\n", conn.RemoteAddr())
				}
				logout(uhash, this.Debug, this.db) //logout(chigh_id, cport, this.Debug, this.db)
			} else if errors.Is(err, net.ErrClosed) {
				if this.Debug {
					fmt.Println("DEBUG: conn closed due to invalid client data")
				}
			}else {
				fmt.Println("ERROR: from read:", err.Error())
			}
			client.DeComp.Close()
			client.Comp.Close()
			client.Conn.Close()
			return
		}
		if this.Debug {
			fmt.Printf("DEBUG: type 0x%02x\n", buf[0])
		}
		if buf[0] == 0x01 {
			uhash = login(buf, protocol, conn, this.Debug, this.db,HighId(this.Host),uint16(this.Port), this.Ssname, this.Ssdesc, this.Ssmsg, this.getTCPFlags())//chigh_id, cport, uhash = login(buf, protocol, conn, this.Debug, this.db)
		} else if buf[0] == 0x14 {
			listservers(buf, protocol, conn, this.Debug, buflen)
		} else if buf[0] == 0x15 {
			offerfiles(buf, protocol, &client, this.Debug, buflen, this.db ,uhash)  //offerfiles(buf, protocol, conn, this.Debug, buflen)
		} else if buf[0] == 0x16 {
			searchfiles(buf, protocol, conn, this.Debug, buflen, this.db)
		} else if buf[0] == 0x19 {
			filesources(buf, uhash, protocol, conn, this.Debug, buflen, this.db)
		} else if buf[0] == 0x1c {
			requestcallback(buf, protocol, conn, this.Debug, buflen)
		} else if buf[0] == 0x9a {
			udpfilesources(buf, protocol, conn, this.Debug, buflen)
		}
	}
}

func (this *SockSrv) yoursam() string {
	return fmt.Sprintf("%s:%d", this.SAM, this.SAMPort)
}

func (this *SockSrv) Start() {
	if this.SQL {
		if this.Debug {
			fmt.Println("With SQL")	
			fmt.Printf("String: %s:%s@tcp(%s:%d)/%s\n", this.SqlUser, this.SqlPW, this.SqlAddr, this.SqlPort, this.SqlDB)
			fmt.Println("SQL DRIVER", this.SqlDriver)
		}

		db, err := sql.Open(this.SqlDriver, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", this.SqlUser, this.SqlPW, this.SqlAddr, this.SqlPort, this.SqlDB))
		if err != nil {
			fmt.Println("ERROR:", err.Error())
			return
		}
		this.db = db
		//res, err := db.Query("select * from clients")
		//if err != nil {
		//	fmt.Println("ERROR:", err.Error())
		//	return
		//}
	        //defer res.Close()
	}
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
	if this.SQL {
		defer this.db.Close()
	}
	defer this.listener.Close()
	return
}

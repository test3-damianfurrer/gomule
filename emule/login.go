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
	"net"
	"database/sql"
)

func login(buf []byte, protocol byte, conn net.Conn, debug bool, db *sql.DB) {
	if debug {
		fmt.Println("DEBUG: Login")
	}
	high_id := highId(conn.RemoteAddr().String())
	//uuidsql := fmt.Sprintf("0x%x",buf[1:17])
	port := byteToInt16(buf[21:23])
	tags := byteToInt32(buf[23:27])
	if debug {
		uuid := fmt.Sprintf("%x-%x-%x-%x-%x-%x-%x-%x",
		buf[1:3], buf[3:5], buf[5:7], buf[7:9], buf[9:11], buf[11:13],
		buf[13:15], buf[15:17])
		fmt.Println("DEBUG: highid:", high_id)
		fmt.Println("DEBUG: uuid:  ", uuid)
		fmt.Println("DEBUG: port:  ", port)
		fmt.Println("DEBUG: tagscount:  ", tags)
		fmt.Println("DEBUG: port bytes:  ", buf[21:23])
		fmt.Println("DEBUG: tagscount bytes:  ", buf[23:27])
		fmt.Println("DEBUG: pre str tag bytes:  ", buf[27:31])
		//fmt.Println("DEBUG: other:  ", buf[27:50])
		//+4 some codes    [2 1 0 1 21 0 104 116
		//21 0 = lenght, 104 116 .. string
		strlen := byteToInt16(buf[31:33])
		str := fmt.Sprintf("%s",buf[33:33+strlen])
		fmt.Println("DEBUG: user name:  ", str)
		fmt.Println("DEBUG: vers tag:  ", buf[33+strlen:33+strlen+8])
		fmt.Println("DEBUG: port tag:  ", buf[33+strlen+8:33+strlen+16])
		fmt.Println("DEBUG: flag tag:  ", buf[33+strlen+16:33+strlen+24])
		//strlen + 3*8bytes should exactly be the end of the buffer //confirmed
	}
	
	//res, err := db.Exec(fmt.Sprintf("INSERT INTO clients(hash, id_ed2k, ipv4, port, online) VALUES (%s,%d, %d, %d, %d)",uuidsql,high_id,high_id,port,1))
	res, err := db.Exec("INSERT INTO clients(hash, id_ed2k, ipv4, port, online) VALUES (?, ?, ?, ?, ?)",buf[1:17],high_id,high_id,port,1)
	fmt.Println("DEBUG: res: ",res)
	fmt.Println("DEBUG: err: ",err)
	if err != nil {
		fmt.Println("ERROR: ",err.Error())
		return
    	}

	data := []byte{protocol,
		8, 0, 0, 0,
		0x38,
		5, 0,
		'h', 'e', 'l', 'l', 'o'}
	if debug {
		fmt.Println("DEBUG: login:", data)
	}
	conn.Write(data)

	data = []byte{protocol,
		9, 0, 0, 0,
		0x40,
		0, 0, 0, 0,
		1, 0, 0, 0}
	high_id_b := uint32ToByte(high_id)
	for i := 0; i < len(high_id_b); i++ {
		data[i+6] = high_id_b[i]
	}
	if debug {
		fmt.Println("DEBUG: login:", data)
	}
	conn.Write(data)
	
	data = []byte{protocol,
		9, 0, 0, 0,
		0x34,       //server status
		1, 0, 0, 0, //user count
		1, 0, 0, 0} //file count
	if debug {
		fmt.Println("DEBUG: login:", data)
	}
	conn.Write(data)
	//0x41 server identification missing
}

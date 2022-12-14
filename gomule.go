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

package main

import (
	"flag"
	"fmt"

	"github.com/test3-damianfurrer/gomule/emule"
)

var (
	debug   bool
	host    string
	port    int
	version bool
	i2p     bool
	sam     string
	samport int
	usesql  bool
	sqldriver string
	sqluser  string
	sqlpw    string
	sqladdr	 string
	sqlport	 int
	sqldb    string
)

func init() {
	flag.BoolVar(&debug, "d", false, "Debug")
	flag.StringVar(&host, "h", "localhost", "Host address")
	flag.IntVar(&port, "p", 7111, "Port number")
	flag.BoolVar(&i2p, "i", false, "Use I2P")
	flag.StringVar(&sam, "s", "127.0.0.1", "SAM host address")
	flag.IntVar(&samport, "sp", 7656, "SAM port number")
	flag.BoolVar(&version, "v", false, "Version")
	flag.BoolVar(&usesql, "us", false, "Use SQL DB")
	flag.StringVar(&sqldriver, "sd", "mysql", "SQL driver")
	flag.StringVar(&sqluser, "su", "user", "SQL user")
	flag.StringVar(&sqlpw, "pw", "password", "SQL password")
	flag.StringVar(&sqldb, "db", "gomule", "SQL DB name")
	flag.StringVar(&sqladdr, "ssi", "127.0.0.1", "SQL Server ip/domain")
	flag.IntVar(&sqlport, "ssp", 3306, "SQL port number")
	

}

func main() {
	flag.Parse()

	if version {
		fmt.Println("GoMule server Version 1.0")
		fmt.Println("Copyright 2013 Leslie Zhai")
		return
	}

	sock := emule.NewSockSrv(host, port, debug)
	
	/*
	SupportGzip
	SupportNewTags
	SupportUnicode
	SupportRelSearch
	SupportTTagInteger
	SupportLargeFiles
	SupportObfuscation
	*/
	sock.SupportGzip=true
	//sock.SupportLargeFiles=true
	
	sock.Ssname = "Test Server"
	sock.Ssdesc = "Gomule a Testing Server"
	sock.Ssmsg = "server version 0.0.1 (gomule)\nwarning - warning you\nHeLlo Brother in christ\n->New Line"
	sock.I2P = i2p
	sock.SAM = sam
	sock.SAMPort = samport
	sock.SQL = usesql
	sock.SqlDriver = sqldriver
	sock.SqlUser = sqluser
	sock.SqlPW = sqlpw
	sock.SqlAddr = sqladdr
	sock.SqlPort = sqlport
	sock.SqlDB = sqldb
	sock.Start()
	defer sock.Stop()
}

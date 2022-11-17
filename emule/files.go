package emule

import (
	"fmt"
	"net"
	libdeflate "github.com/4kills/go-libdeflate/v2"
)
func prconefile(filehashbuf []byte, filename string, fsize uint32, filetype string, debug bool){
	fuuid := fmt.Sprintf("%x-%x-%x-%x-%x-%x-%x-%x",
		filehashbuf[0:2], filehashbuf[2:4], 
		filehashbuf[4:6], filehashbuf[6:8],
		filehashbuf[8:10], filehashbuf[10:12], 
		filehashbuf[12:14], filehashbuf[14:16])
	if debug {
    		fmt.Println("DEBUG: File hash:", fuuid)  
    		fmt.Println("DEBUG: File name:", filename)
    		fmt.Println("DEBUG: File type:", filetype)
		fmt.Println("DEBUG: File size:", fsize)
	}
}

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
  iteration := 0
  byteoffset := uint32(4)
  //debugloop:=debug
  debugloop:=false
  for{
    //if byteoffset >= uint32(blen) {
    if iteration >= 200{
	    if debug {	
		    fmt.Println("DEBUG: exiting, byteoffset >= bufferlength", blen)
		    fmt.Println("byteoffset", byteoffset)
		    fmt.Println("iteration", iteration)
	    }
	    break;
    }
    if debugloop {
      fmt.Println("DEBUG: byteoffset", byteoffset)
      fmt.Println("DEBUG: iteration", iteration)
    }
    filehashbuf := buf[byteoffset+0:byteoffset+16]
    
	 //obfuscated
    //fmt.Println("DEBUG: client id:", buf[byteoffset+16:byteoffset+20])
    //fmt.Println("DEBUG: client port:", buf[byteoffset+20:byteoffset+22])
    itag := byteToInt32(buf[byteoffset+22:byteoffset+26])
    if debugloop {
    	fmt.Println("DEBUG: 1. tag count:", itag)
    }
	  //skip 4 [2 1 0 1] 
    strlen := uint32(byteToInt16(buf[byteoffset+30:byteoffset+32]))
    strbuf := buf[byteoffset+32:byteoffset+32+strlen]
    fname := fmt.Sprintf("%s",strbuf)
    //[3 1 0 2]
    fsize := byteToUint32(buf[byteoffset+32+strlen+4:byteoffset+32+strlen+8])
   
    if itag > 2 {
    	//[2 1 0 3]
	strlentype := uint32(byteToInt16(buf[byteoffset+32+strlen+12:byteoffset+32+strlen+14]))
    	strbuf = buf[byteoffset+32+strlen+14:byteoffset+32+strlen+14+strlentype]
    	//str = fmt.Sprintf("%s",strbuf)
    	
    	prconefile(filehashbuf, fname, fsize, fmt.Sprintf("%s",strbuf), debugloop)
    	byteoffset = byteoffset+32+strlen+14+strlentype
	    //in theory needs to be able to handle more tags
    } else {
    	prconefile(filehashbuf, fname, fsize, "", debugloop)
	byteoffset = byteoffset+32+strlen+8
    }
    //fmt.Println("DEBUG: 30 bytes more:", buf[byteoffset+36+strlen+14+strlentype:byteoffset+36+strlen+14+strlentype+30])
    iteration+=1
	  
    if debugloop {
      fmt.Println("DEBUG: new byteoffset", byteoffset)
      fmt.Println("DEBUG: next iteration", iteration)
    }
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
	if debug {
  		fmt.Println("DEBUG: decompressing")
	}
	//blen, decompressed, err = dc.Decompress(bufcomp, nil, 1)
  	blen, decompressed, err = dc.Decompress(bufcomp, nil, 1)
	dc.Close()
	if debug {
		fmt.Println("DEBUG: after decompressing")
	}
	if err != nil {
		fmt.Println("ERROR decompress:", err.Error())
		fmt.Println("ERROR: uncompressed len", blen)
	  	fmt.Println("ERROR: uncompressed buf 10", decompressed[0:10])
		return
	}
  	fmt.Println("DEBUG: uncompressing processed bytes", blen)
  	//fmt.Println("DEBUG: uncompressed buf 10", decompressed[blen+0:blen+10])
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


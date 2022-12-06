package emule

import (
	"fmt"
	//"net"
	//"database/sql"
  	"strings"
)

//Field   Type    Null    Key     Default Extra
//id      bigint unsigned NO      PRI     NULL    auto_increment
//file_hash       binary(16)      NO      MUL     0x30
//user_hash       binary(16)      NO              0x30
//name    varchar(255)    NO      MUL
//ext     varchar(8)      NO
//time_offer      timestamp       NO              CURRENT_TIMESTAMP       DEFAULT_GENERATED
//type    enum('Image','Audio','Video','Pro','Doc','')    NO
//rating  tinyint unsigned        NO              0
//title   varchar(128)    NO
//artist  varchar(128)    NO
//album   varchar(128)    NO
//length  int unsigned    NO              0
//bitrate int unsigned    NO              0
//codec   varchar(32)     NO
//online  tinyint(1)      NO              0
//complete        tinyint(1)      NO              0


func search2query(search string)(sqlquery string, strarr []string){
  sqlquery = "select name, ext, type, rating, file_hash from sources WHERE "
  strarr = strings.Split(search," ")
  for i := 0; i < len(strarr); i++ {
	  strarr[i] = "%"+strarr[i]+"%"
	  if i < len(strarr)-1 {
	  	sqlquery += "name like ? AND "
	  } else {
		sqlquery += "name like ?"
	  }
	  fmt.Println("String: ",i,strarr[i])
  }
  fmt.Println("query: ",strarr)
  return
}

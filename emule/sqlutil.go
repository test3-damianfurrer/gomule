package emule

import (
	"fmt"
	//"net"
	//"database/sql"
  	"strings"
)

func search2query(search string)(sqlquery string, strarr []string){
  sqlquery = "select name, ext, type, rating from sources WHERE "
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

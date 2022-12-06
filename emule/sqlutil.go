package emule

import (
	"fmt"
	//"net"
	//"database/sql"
  	"strings"
)

func search2query(search string)(sqlquery string){
  sqlquery = "select name, ext, type, rating from sources WHERE "
  strarr := strings.Split(search," ")
  for i := 0; i < len(strarr); i++ {
	  strarr[i] = "%"+strarr[i]+"%"
	  sqlquery += "name like ? AND "
	  fmt.Println("String: ",i,strarr[i])
  }
  fmt.Println("query: ",strarr)
  return
}

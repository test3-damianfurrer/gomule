package emule

import (
	"fmt"
	//"net"
	//"database/sql"
  	"strings"
)

func search2query(search string)(sqlquery string){
  //sqlquery = "test"
  strarr := strings.Split(search," ")
  fmt.Println("query: ",strarr)
  return
}

package emule

import (
	"fmt"
	"net"
	"database/sql"
  	"strings"
)

func search2query(string search)(string sqlquery){
  //sqlquery = "test"
  strarr := strings.Split(search)
  fmt.Println("query: ",strarr)
  return
}

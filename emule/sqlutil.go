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
  for i := 0; i < len(strarr); i++ {
	  fmt.Println("String: ",i,strarr[i])
  }
  fmt.Println("query: ",strarr)
  return
}

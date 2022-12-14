package emule

import (
	"fmt"
	//"net"
	"database/sql"
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
func filename2ext(filename string) string {
	strarr := strings.Split(filename,".")
	sindex := len(strarr)-1
	if sindex == 0 {
		return ""
	}
	slen := len(strarr[sindex])
	if slen > 8 {
		return ""
	}
	return strings.ToUpper(strarr[sindex])
}

func stringifyConstraint(in *Constraint, params *[]interface{})(ret string){
	switch in.Type {
		case C_AND:
			ret = "("+stringifyConstraint(in.Left,params)+") AND ("+stringifyConstraint(in.Right,params)+")"
		case C_OR:
			ret = "("+stringifyConstraint(in.Left,params)+") OR ("+stringifyConstraint(in.Right,params)+")"
		case C_NOT:
			ret = "("+stringifyConstraint(in.Left,params)+") NOT ("+stringifyConstraint(in.Right,params)+")"
		case C_MAIN:
			strarr := strings.Split(fmt.Sprintf("%s",in.Value)," ")
			ret = ""
  			for i := 0; i < len(strarr); i++ {
				if i != 0 {
					ret += " AND "
				}
				ret += "sources.name like ?"
				*params = append(*params,"%"+strarr[i]+"%")
			}
		case C_CODEC:
		case C_MINSIZE:
			*params = append(*params,ByteToUint32(in.Value))
			fmt.Println("DEBUG(sqlutil.go): minsize ",ByteToUint32(in.Value),in.Value)
			ret = "files.size >= ?"
		case C_MAXSIZE:
			*params = append(*params,ByteToUint32(in.Value))
			fmt.Println("DEBUG(sqlutil.go): maxsize ",ByteToUint32(in.Value),in.Value)
			ret = "files.size <= ?"
		case C_FILETYPE:
			*params = append(*params,fmt.Sprintf("%s",in.Value))
			ret = "sources.type = ?"
		case C_FILEEXT:
			*params = append(*params,fmt.Sprintf("%s",in.Value))
			ret = "sources.ext like ?"
		default:
			fmt.Println("ERROR: undefined Constraint Type", in.Type)
	}
	return
}
func constraintsearch2query(in *Constraint, params *[]interface{})(sqlquery string){
	sqlquery = "select sources.name, sources.ext, sources.type, sources.rating, sources.file_hash, files.size from sources left join files on sources.file_hash=files.hash WHERE "
	sqlquery += stringifyConstraint(in, params)
	return
}

func search2query(search string)(sqlquery string, strarr []string){
  sqlquery = "select sources.name, sources.ext, sources.type, sources.rating, sources.file_hash, files.size from sources left join files on sources.file_hash=files.hash WHERE "
  strarr = strings.Split(search," ")
  for i := 0; i < len(strarr); i++ {
	  strarr[i] = "%"+strarr[i]+"%"
	  if i < len(strarr)-1 {
	  	sqlquery += "sources.name like ? AND "
	  } else {
		sqlquery += "sources.name like ?"
	  }
	  fmt.Println("String: ",i,strarr[i])
  }
  fmt.Println("query: ",strarr)
  return
}

func readRowUint32(query string,db *sql.DB) uint32 {
	var value uint32
	err := db.QueryRow(query).Scan(&value)
    	if err != nil {
		fmt.Println("ERROR(readRowUint32): ",err.Error())
	}
	return value
}

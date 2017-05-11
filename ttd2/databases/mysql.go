package databases

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/msfy/fengyun.cc/home/utils"
)

var Mysql *sql.DB

func init() {
    var err error
    const dataSourceName = "root:Tczkj110@tcp(rm-2ze8e8qzj6aa54837o.mysql.rds.aliyuncs.com:3306)/zgamedaer?charset=utf8"
    Mysql, err = sql.Open("mysql", dataSourceName)
    utils.CheckError(err)
}

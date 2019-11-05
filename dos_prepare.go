package main

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db      *sql.DB
	preStmt *sql.Stmt
	err     error
)

func main() {
	db,err = sql.Open("mysql", "root:123456@tcp(10.179.132.206:3306)/test?charset=utf8")
	defer db.Close()
	preStmt, err = db.Prepare("insert into testprepare (`i1`) values (?)")
	defer preStmt.Close()
        stmt.Exec(20)
 } 




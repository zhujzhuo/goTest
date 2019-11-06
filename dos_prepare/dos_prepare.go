package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/logger"
	"log"
	"os"
	"time"
)

var (
	db   *sql.DB
	stmt *sql.Stmt
	err  error
)
var (
	id int
	name string
)

const logPath = "./dos_prepare.log"

var verbose = flag.Bool("verbose", false, "print info level logs to stdout") //控制是否将info日志输出到屏幕上

func main() {
	fmt.Printf("start main func at: %v\n", time.Now().Format("2006-1-2 15:04:05"))
	flag.Parse()
	lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	defer lf.Close()

	defer logger.Init("dos_prepare", *verbose, true, lf).Close()
	logger.SetFlags(log.LstdFlags)

	logger.Info("I'm about to do something!")

	db, err = sql.Open("mysql", "mysqltest:123456@tcp(localhost:3306)/test?charset=utf8mb4")
	if err != nil {
		logger.Fatalf("Failed to open mysql connection: %v", err)
	}
	logger.Info("just test for  logger")

	defer db.Close()
	stmt, err = db.Prepare("insert into testprepare (i1) values (?)")
	defer stmt.Close()
	res, err := stmt.Exec(20)   //写入使用Exec
        //rows, err := stmt.Query(1)   //查询使用Query
	lastId, err := res.LastInsertId()
	rowCnt, err := res.RowsAffected()
	logger.Infof("ID = %d, affected = %d\n", lastId, rowCnt)

	//rows, err := db.Query("select id, name from users where id = ?", 1)
	rows, err := db.Query("select id, name from users")
	if err != nil {
		logger.Fatalf("query error :%v", err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			logger.Fatalf("",err)
		}
		fmt.Println("id:",id,"name:",name)
	}
	err = rows.Err()
	if err != nil {
		logger.Fatalf("",err)
	}
        fmt.Println()
}

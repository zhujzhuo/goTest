package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/logger"
)

var (
	db   *sql.DB
	stmt *sql.Stmt
	err  error
)
var (
	id   int
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

	//db, err = sql.Open("mysql", "mysqltest:123456@tcp(localhost:3306)/test?charset=utf8mb4")
	db, err = sql.Open("mysql", "root:123456@tcp(10.179.132.206:3306)/test?charset=utf8mb4")
	if err != nil {
		logger.Fatalf("Failed to open mysql connection: %v", err)
	}
	logger.Info("just test for  logger")

	defer db.Close()
	//test  prepare
	stmt, err = db.Prepare("insert into testprepare (i1) values (?)")
	defer stmt.Close()
	res, err := stmt.Exec(20) //写入使用Exec
	//rows, err := stmt.Query(1)   //查询使用Query
	lastId, err := res.LastInsertId()
	rowCnt, err := res.RowsAffected()
	logger.Infof("ID = %d, affected = %d\n", lastId, rowCnt)

	//test select query
	//rows, err := db.Query("select id, name from users where id = ?", 1)
	rows, err := db.Query("select id, name from users")
	//_, err = db.Exec("Update  users set  name='update' where name='Dolly'")  //直接执行更新
	//_, err = db.Exec("DELETE FROM users where id=5")   //直接执行删除

	//捕获特殊的报错信息,下面的不是常规的写法
	//
	//if strings.Contains(err.Error(), "Access denied") {
	//Handle the permission-denied error
	//}
	/* 如下是捕获异常的正确方法
	        if driverErr, ok := err.(*mysql.MySQLError); ok { // Now the error number is accessible directly
		if driverErr.Number == 1045 {
			// Handle the permission-denied error
		}
	        }
	*/
	/*
	   if driverErr, ok := err.(*mysql.MySQLError); ok {
	   	if driverErr.Number == mysqlerr.ER_ACCESS_DENIED_ERROR {
	   		// Handle the permission-denied error
	   	}
	   }
	*/
	if err != nil {
		logger.Fatalf("query error :%v", err)
	}
	defer rows.Close()
	for rows.Next() { //遍历结果
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("id:", id, "name:", name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	//test delete insert  update  for exec
	stmt, err = db.Prepare("INSERT INTO users(name) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err = stmt.Exec("Dolly")
	if err != nil {
		log.Fatal(err)
	}
	lastId, err = res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err = res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

	//1.  prepare创建链接后，放回连接池【与分片绑定】；exec再从连接池取出【可能是其他连接】
	//2.  直接exec，创建/获取连接，执行；使用同一个连接
	// transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("INSERT INTO users(name) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // danger!
	for i := 0; i < 10; i++ {
		_, err = stmt.Exec(i)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	// stmt.Close() runs here!
	fmt.Println()
}

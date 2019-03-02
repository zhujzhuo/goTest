package main
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	_ "net/http/pprof"
	"net/http"
	"time"
)
var	(
	db	*sql.DB
	preStmt *sql.Stmt
	err	error
)
//mysql -utestomega_rw  -h10.94.97.147 -P4001  -p'TestOmegaRW@0612'
func main() {
	go func() {
		http.ListenAndServe("localhost:8888", nil)
	}()
        db,err = sql.Open("mysql", "testomega_rw:TestOmegaRW@0612@tcp(10.94.97.147:4001)/db_omega_web?charset=utf8&collation=utf8_general_ci")
	//db,err = sql.Open("mysql", "user1:password1@tcp(10.94.99.67:3306)/db_omega_web?charset=utf8&collation=utf8_general_ci")
	db.SetMaxOpenConns(2)
	db.SetMaxIdleConns(2)
	if err != nil {
		fmt.Printf("db open err:%v",err)
		err = nil
	}
	defer db.Close()
	preStmt, err = getStat()
	err = nil
	defer preStmt.Close()
	fmt.Printf("get here, start test\n")
	test()
}

func getStat()(*sql.Stmt, error){
	sqlById := "insert into t_crash_analysis_monthly (`err_tag`) values (?);"
	preStmt, err = db.Prepare(sqlById)
	if err != nil{
		fmt.Printf("this is getStat func err:%v\n", err)
		err = nil
	}
	return preStmt, err
}

func test(){
	fmt.Print("enter test\n")
	wg := &(sync.WaitGroup{})
	for i:=0; i < 20000;i ++{
		wg.Add(1)
		go testDb(wg, i)
		time.Sleep(1e9)
		//go testDb(wg, i)
		//wg.Add(1)
		//go testDb2(wg, i)
	}
	//wg.Add(1)
	//go testDb(wg, 0)

	wg.Wait()
}

func testDb(wg *sync.WaitGroup, no int){
	fmt.Printf("this is %dth func\n", no)
	begTime := time.Now()
	defer wg.Done()
	selectAll := "select id from t_crash_analysis_monthly where 1=1"
	rowsAll,err := db.Query(selectAll)
	defer rowsAll.Close()
	if err!= nil{
		fmt.Printf("selectAll err is:%v, this is %dth func\n", err.Error(), no)
		err = nil
	}
	var (
		id int64
	)
	count := 0
	for rowsAll.Next(){
		rowsAll.Scan(&id)
		updateOne := "update t_crash_analysis_monthly set `err_tag`='test_err_tag' where id=?"
		_, err = db.Exec(updateOne, id)
		if err != nil{
			fmt.Printf("db err is:%v\n", err)
			err = nil
		}
		count ++
		//fmt.Printf("this is %dth func|testDb|get message:%d,total count is:%d\n",no, id, count)
		//time.Sleep(1e8)
	}
	endTime := time.Now()
	fmt.Printf("this is %dth func;spend %d millisecond\n", no, endTime.Sub(begTime)/time.Millisecond)
	fmt.Printf("count is %v\n",count)
	//time.Sleep(1e12)
	return
}

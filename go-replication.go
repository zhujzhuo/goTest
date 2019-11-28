package main

//这个binlog解析和Mysql版本有关系，8.0的版本暂时不支持
import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
)

func main() {
	// Create a binlog syncer with a unique server id, the server id must be different from other MySQL's.
	// flavor is mysql or mariadb
	cfg := replication.BinlogSyncerConfig{
		ServerID: 100,
		Flavor:   "mysql",
		Host:     "10.179.132.206",
		Port:     3306,
		User:     "root",
		Password: "123456",
	}
	syncer := replication.NewBinlogSyncer(cfg)

	// binlogFile := "binlog.000005"
	// var binlogPos uint32 = 4
	// Start sync with specified binlog file and position
	streamer, _ := syncer.StartSync(mysql.Position{"dd-bin.000079", 4})
	// or you can start a gtid replication like
	// streamer, _ := syncer.StartSyncGTID(gtidSet)
	// the mysql GTID set likes this "de278ad0-2106-11e4-9f8e-6edd0ca20947:1-2"
	// the mariadb GTID set likes this "0-1-100"
	fmt.Println(time.Now())
	//for {
        //	ev, err := streamer.GetEvent(context.Background())
	//	// Dump event
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	ev.Dump(os.Stdout)
	//}
	// or we can use a timeout context
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		ev, err := streamer.GetEvent(ctx)
		cancel()

		if err == context.DeadlineExceeded {
			// meet timeout
			continue
		}
		ev.Dump(os.Stdout)
	}
}

/*
[2019/11/09 22:05:31] [info] binlogsyncer.go:133 create BinlogSyncer with config {100 mysql 10.179.132.206 3306 root    false false <nil> false UTC false 0 0s 0s 0 false}
[2019/11/09 22:05:31] [info] binlogsyncer.go:354 begin to sync binlog from position (dd-bin.000079, 4)
[2019/11/09 22:05:31] [info] binlogsyncer.go:203 register slave for master server 10.179.132.206:3306
2019-11-09 22:05:31.220453 +0800 CST m=+0.044451044
=== RotateEvent ===
Date: 1970-01-01 08:00:00
Log position: 0
Event size: 40
Position: 4
Next log name: dd-bin.000079

=== FormatDescriptionEvent ===
Date: 2019-11-09 22:04:29
Log position: 123
Event size: 119
Version: 4
Server version: 5.7.19-log
Checksum algorithm: 1

=== PreviousGTIDsEvent ===
Date: 2019-11-09 22:04:29
Log position: 194
Event size: 71
Event data:
00000000  01 00 00 00 00 00 00 00  73 a4 15 36 01 b1 11 e9  |........s..6....|
00000010  84 eb ca fd d0 1d 4b b9  01 00 00 00 00 00 00 00  |......K.........|
00000020  01 00 00 00 00 00 00 00  c6 97 28 00 00 00 00 00  |..........(.....|

=== GTIDEvent ===
Date: 2019-11-09 22:05:08
Log position: 259
Event size: 65
Commit flag: 0
GTID_NEXT: 73a41536-01b1-11e9-84eb-cafdd01d4bb9:2660294
LAST_COMMITTED: 0
SEQUENCE_NUMBER: 1

=== QueryEvent ===
Date: 2019-11-09 22:05:08
Log position: 331
Event size: 72
Slave proxy ID: 730
Execution time: 0
Error code: 0
Schema: test
Query: BEGIN

=== RowsQueryEvent ===
Date: 2019-11-09 22:05:08
Log position: 419
Event size: 88
Query: insert into  t select null,'北京','zhujzhuo',20,NULL,NULL,NULL

=== TableMapEvent ===
Date: 2019-11-09 22:05:08
Log position: 479
Event size: 60
TableID: 367
TableID size: 6
Flags: 1
Schema: test
Table: t
Column count: 7
Column type:
00000000  03 0f 0f 03 0f 0f 0f                              |.......|
NULL bitmap:
00000000  70                                                |p|

=== WriteRowsEventV2 ===
Date: 2019-11-09 22:05:08
Log position: 539
Event size: 60
TableID: 367
Flags: 1
Column count: 7
Values:
--
0:206
1:"北京"
2:"zhujzhuo"
3:20
4:<nil>
5:<nil>
6:<nil>

=== XIDEvent ===
Date: 2019-11-09 22:05:08
Log position: 570
Event size: 31
XID: 2937

[2019/11/09 22:05:31] [info] binlogsyncer.go:723 rotate to (dd-bin.000079, 4)
*/

package main

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
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "mysqltest",
		Password: "123456",
	}
	syncer := replication.NewBinlogSyncer(cfg)

	// binlogFile := "binlog.000005"
	// var binlogPos uint32 = 4
	// Start sync with specified binlog file and position
	streamer, _ := syncer.StartSync(mysql.Position{"binlog.000005", 4})
	// or you can start a gtid replication like
	// streamer, _ := syncer.StartSyncGTID(gtidSet)
	// the mysql GTID set likes this "de278ad0-2106-11e4-9f8e-6edd0ca20947:1-2"
	// the mariadb GTID set likes this "0-1-100"
	fmt.Println(time.Now())
	for {
		ev, err := streamer.GetEvent(context.Background())
		// Dump event
		if err != nil {
			fmt.Println(err)
		}
		ev.Dump(os.Stdout)
	}
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

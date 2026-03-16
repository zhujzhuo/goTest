package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/replication"
)

func main() {
	cfg := replication.BinlogSyncerConfig{
		ServerID: 100,
		Flavor:   "mysql",
		Host:     "192.168.215.2",
		Port:     3306,
		User:     "cruiser",
		Password: "cruiser@123",
	}

	syncer := replication.NewBinlogSyncer(cfg)
	streamer, _ := syncer.StartSync(mysql.Position{"mysql-bin.000007", 4})
	fmt.Println(time.Now())
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		ev, err := streamer.GetEvent(ctx)
		cancel()
		if err == context.DeadlineExceeded {
			continue
		}
		if rowsEvent, ok := ev.Event.(*replication.RowsEvent); ok {
			fmt.Println(ev.Header.EventType.String())
			for i, row := range rowsEvent.Rows {
				fmt.Println("row:\n", i, row)
			}
		}
		if rotateEvent, ok := ev.Event.(*replication.RotateEvent); ok {
			fmt.Println(ev.Header.EventType.String())
			fmt.Println(rotateEvent.Position)
			ev.Dump(os.Stdout)
		}
	}
}

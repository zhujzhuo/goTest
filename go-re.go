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
	cfg := replication.BinlogSyncerConfig{
		ServerID: 100,
		Flavor:   "mysql",
		Host:     "10.179.132.206",
		Port:     3306,
		User:     "root",
		Password: "123456",
	}

	syncer := replication.NewBinlogSyncer(cfg)
	streamer, _ := syncer.StartSync(mysql.Position{"dd-bin.000079", 4})
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
				fmt.Println("row:\n", i,row)
			}
		}
		if rotateEvent, ok := ev.Event.(*replication.RotateEvent); ok {
                        fmt.Println(ev.Header.EventType.String())
                        fmt.Println(rotateEvent.Position)
		        ev.Dump(os.Stdout)
		}
	}
}


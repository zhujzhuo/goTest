package main

//https://github.com/google/logger
//使用 google logger 和自定义error
import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/logger"
)

const logPath = "./google_logger.log"

var verbose = flag.Bool("verbose", false, "print info level logs to stdout")

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	// return fmt.Sprintf("at %v, %s",
	// 	e.When, e.What)
	return fmt.Sprintf("some errors")
}
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	flag.Parse()

	lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	defer lf.Close()

	defer logger.Init("LoggerExample", *verbose, true, lf).Close()
	logger.SetFlags(log.LstdFlags)

	logger.Info("I'm about to do something!")
	if err := run(); err != nil {
		logger.Errorf("Error running doSomething: %v", err)
	}
}

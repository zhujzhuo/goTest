package main

import (
	"encoding/json"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// logpath 日志文件路径
// loglevel 日志级别
func initLogger(logpath string, loglevel string) *zap.Logger {

	hook := lumberjack.Logger{
		Filename:   logpath, // 日志文件路径
		MaxSize:    128,     // megabytes
		MaxBackups: 30,      // 最多保留30个备份
		MaxAge:     7,       // days
		Compress:   true,    // 是否压缩 disabled by default
	}

	w := zapcore.AddSync(&hook)

	// 设置日志级别,debug可以打印出info,debug,warn；info级别可以打印warn，info；warn只能打印warn
	// debug->info->warn->error
	var level zapcore.Level
	switch loglevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	// 时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		w,
		level,
	)

	logger := zap.New(core,zap.AddCaller(),zap.AddCallerSkip(1))
	logger.Info("DefaultLogger init success")
	return logger
}

type Test struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	t := &Test{
		Name: "xiaoming",
		Age:  12,
	}
	data, err := json.Marshal(t)
	fmt.Println(data)
	if err != nil {
		fmt.Println("marshal is failed,err: ", err)
	}

	// 历史记录日志名字为：all-2018-11-15T07-45-51.763.log，服务重新启动，日志会追加，不会删除
	logger := initLogger("./all.log", "debug")
	for i := 0; i < 6; i++ {
		logger.Info(fmt.Sprint("test log ", i), zap.Int("line", 47))
		logger.Debug(fmt.Sprint("debug log ", i), zap.ByteString("level", data))
		logger.Info(fmt.Sprint("Info log ", i), zap.String("level", `{"a":"4","b":"5"}`))
		logger.Warn(fmt.Sprint("Info log ", i), zap.String("level", `{"a":"7","b":"8"}`))
	}
/*
2019-03-21T15:47:24.214+0800    INFO    goTest/zap_exercise.go:74       test log 5      {"line": 47}
2019-03-21T15:47:24.214+0800    DEBUG   goTest/zap_exercise.go:75       debug log 5     {"level": "{\"name\":\"xiaoming\",\"age\":12}"}
2019-03-21T15:47:24.214+0800    INFO    goTest/zap_exercise.go:76       Info log 5      {"level": "{\"a\":\"4\",\"b\":\"5\"}"}
2019-03-21T15:47:24.214+0800    WARN    goTest/zap_exercise.go:77       Info log 5      {"level": "{\"a\":\"7\",\"b\":\"8\"}"}
*/
}


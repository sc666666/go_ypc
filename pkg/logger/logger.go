package logger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go_ypc/pkg/config"
	"io"
	"log"
	"os"
	"time"
)

var (
	sugarLogger *zap.SugaredLogger

	// 日志文件配置
	logFilePath   string // 日志文件路径
	logNamePrefix string // 日志文件名称前缀
	logNameFormat string // 日志文件名称格式
	logFileExt    string // 日志文件扩展名称

	// 日志切割归档配置
	MaxSize    int  // 在进行切割之前，日志文件的最大大小（以MB为单位）
	MaxBackups int  // 保留旧文件的最大个数
	MaxAge     int  // 保留旧文件的最大天数
	Compress   bool // 是否压缩/归档旧文件

	hook io.Writer
)

// 初始化 zap
func InitLogger(jsonFormat bool, logInConsole bool, showLine bool) {
	logFilePath = config.GetString("logging.channels." + config.GetString("logging.default") + ".path")
	logNamePrefix = config.GetString("app.name")
	logNameFormat = config.GetString("logging.channels." + config.GetString("logging.default") + ".format")
	logFileExt = config.GetString("logging.channels." + config.GetString("logging.default") + ".ext")

	MaxSize = config.GetInt("logging.cutting.max_size")
	MaxBackups = config.GetInt("logging.cutting.max_backups")
	MaxAge = config.GetInt("logging.cutting.max_age")
	Compress = config.GetBool("logging.cutting.compress")

	// 自定义 zap 配置
	core := zapcore.NewCore(
		getEncoder(jsonFormat),
		getWriteSyncer(logInConsole),
		getLevelEnabler("debug"),
	)

	l := zap.New(core)
	defer l.Sync() // flushes buffer, if any

	if showLine {
		l = l.WithOptions(zap.AddCaller())
	}

	sugarLogger = l.Sugar()
}

// 日志文件格式配置 json/console
func getEncoder(jsonFormat bool) zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:   "msg",
		LevelKey:     "level",
		EncodeLevel:  zapcore.LowercaseLevelEncoder, // 小写编码器 info
		TimeKey:      "time",
		EncodeTime:   zapcore.ISO8601TimeEncoder, // ISO8601 UTC 时间格式
		CallerKey:    "line",
		EncodeCaller: zapcore.ShortCallerEncoder, // 调用者
	}
	if jsonFormat {
		return zapcore.NewJSONEncoder(encoderConfig)
	} else {
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
}

// 日志是否在控制台输出以及是否切割
func getWriteSyncer(logInConsole bool) zapcore.WriteSyncer {
	if true {
		hook = &lumberjack.Logger{
			Filename:   getLogFileFullPath(),
			MaxSize:    MaxSize,
			MaxBackups: MaxBackups,
			MaxAge:     MaxAge,
			Compress:   Compress,
		}
	} else {
		hook = openLogFile()
	}

	if logInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook))
	} else {
		return zapcore.AddSync(hook)
	}
}

// 日志等级
func getLevelEnabler(logLevel string) zapcore.Level {
	switch logLevel {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

// 打开日志文件
func openLogFile() *os.File {
	_, err := os.Stat(logFilePath)
	switch {
	case os.IsNotExist(err):
		makeLogDir(logFilePath)
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}
	logFileFullPath := getLogFileFullPath()
	logFile, err := os.OpenFile(logFileFullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	return logFile
}

// 获取日志文件全路径
func getLogFileFullPath() string {
	return fmt.Sprintf("%s/%s-%s.%s", logFilePath, logNamePrefix, time.Now().Format(logNameFormat), logFileExt)
}

// 创建日志文件夹
func makeLogDir(logFilePath string) {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+logFilePath, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func Debug(args ...interface{}) {
	sugarLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	sugarLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	sugarLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	sugarLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	sugarLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	sugarLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	sugarLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	sugarLogger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	sugarLogger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	sugarLogger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	sugarLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	sugarLogger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	sugarLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	sugarLogger.Fatalf(template, args...)
}

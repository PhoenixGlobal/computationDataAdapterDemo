package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
	"path/filepath"
	"time"
)

func NewLogger(fileName string, logLevel string, isCaller bool, isJson bool) *zap.Logger {
	//fmt.Println("-------------",logLevel)
	var Logger *zap.Logger

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	//log file
	filePath := path.Join(dir, "logs", fileName)

	hook := lumberjack.Logger{
		Filename:   filePath, // path
		MaxSize:    100,      // MaxSize
		MaxBackups: 60,       //
		MaxAge:     7,        //
		Compress:   true,     //
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "linenum",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		//LineEnding:    zapcore.DefaultLineEnding,
		LineEnding:  "\n\n",
		EncodeLevel: zapcore.CapitalLevelEncoder, //
		//EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.In(time.UTC).Format("2006-01-02 15:04:05.999"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		//EncodeCaller:   zapcore.FullCallerEncoder,      //
		EncodeCaller: func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(" " + caller.String() + " ")
		}, //
		EncodeName: zapcore.FullNameEncoder,
	}

	atomicLevel := zap.NewAtomicLevel()

	caller := zap.AddCaller()

	development := zap.Development()
	// filed := zap.Fields(zap.String("serviceName", "serviceName"))
	var enc zapcore.Encoder

	if isJson {
		enc = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		enc = zapcore.NewConsoleEncoder(encoderConfig)
	}
	if logLevel == "debug" {
		atomicLevel.SetLevel(zap.DebugLevel)
		core := zapcore.NewCore(
			enc,
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
			atomicLevel,
		)
		if isCaller {
			Logger = zap.New(core, development, caller)
		} else {
			Logger = zap.New(core, development)
		}
	} else {
		atomicLevel.SetLevel(zap.InfoLevel)

		core := zapcore.NewCore(
			enc,
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
			atomicLevel,
		)
		if isCaller {
			Logger = zap.New(core, development, caller)
		} else {
			Logger = zap.New(core, development)
		}
	}
	return Logger
}

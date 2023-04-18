package log

import (
	"gin-blog/utils"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Lg *zap.Logger

// InitLogger 初始化Logger
func InitLogger() (err error) {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(utils.Level))
	if err != nil {
		return
	}
	var core zapcore.Core
	if utils.AppMode == "debug" {

		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)

	} else {

		core = zapcore.NewCore(encoder, writeSyncer, l)
	}

	Lg = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(Lg) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   utils.Filename,
		MaxSize:    utils.MaxSize,
		MaxBackups: utils.MaxBackups,
		MaxAge:     utils.MaxAge,
		Compress:   utils.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

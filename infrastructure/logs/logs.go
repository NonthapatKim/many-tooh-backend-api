package logs

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.SugaredLogger
var middlewareLogger *zap.SugaredLogger

func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = CustomTimeEncoder

	var err error
	logger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	log = logger.Sugar()

	middlewareLogger = GetMiddleWareLogger()
}

func GetMiddleWareLogger() *zap.SugaredLogger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = CustomTimeEncoder
	config.DisableCaller = true

	logger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	return logger.Sugar()
}

func Info(message string, fields ...interface{}) {
	log.Infof(message, fields...)
}

func LogMiddleWareInfo(message string, fields ...interface{}) {
	middlewareLogger.Infof(message, fields...)
}

func Warn(message string, fields ...interface{}) {
	log.Warnf(message, fields...)
}

func Error(message interface{}, fields ...interface{}) {
	switch v := message.(type) {
	case error:
		log.Errorf(v.Error(), fields...)
	case string:
		log.Errorf(v, fields...)
	}
}

func Errorf(format string, fields ...interface{}) {
	log.Errorf(fmt.Sprintf(format, fields...))
}

func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	utc7 := t.In(time.FixedZone("UTC+7", 7*60*60))
	enc.AppendString(utc7.Format(time.RFC3339))
}

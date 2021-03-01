package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type zapAdapter struct {
	Path        string
	FileName    string
	Level       string // 日志输出的级别
	LogCategory string // 日志类别
	MaxSize     int    //日志文件大小的最大值，单位(M)
	MaxBackups  int    //最多保留备份数
	MaxAge      int    //日志文件保存的时间，单位(天)
	Compress    bool   //是否压缩
	Caller      bool   // 日志是否需要显示调用位置

	logger *zap.Logger
	sugar  *zap.SugaredLogger
}

func NewZapAdapter(path, fileName, Level, logCategory string) *zapAdapter {
	return &zapAdapter{
		Path:        path,
		FileName:    fileName,
		Level:       Level,
		LogCategory: logCategory,
		MaxSize:     1024,
		MaxBackups:  3,
		MaxAge:      7,
		Compress:    true,
	}
}

func (adapter *zapAdapter) Build() {
	writeSyncer := adapter.getLogWriter()
	encoder := adapter.getEncoder()
	var level zapcore.Level
	switch adapter.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "panic":
		level = zap.PanicLevel
	default:
		level = zap.InfoLevel
	}
	core := zapcore.NewCore(encoder, writeSyncer, level)

	//adapter.logger = zap.New(core, zap.AddCaller())
	//adapter.sugar = adapter.logger.Sugar()
	adapter.logger = zap.New(core)
	if adapter.Caller {
		adapter.logger = adapter.logger.WithOptions(zap.AddCaller(), zap.AddCallerSkip(2))
	}
	adapter.sugar = adapter.logger.Sugar().Named(adapter.LogCategory)
}

func (adapter *zapAdapter) getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", adapter.Path, adapter.FileName),
		MaxSize:    adapter.MaxSize,
		MaxBackups: adapter.MaxBackups,
		MaxAge:     adapter.MaxAge,
		Compress:   adapter.Compress,
	}
	//return zapcore.AddSync(lumberJackLogger)
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
}

func (adapter *zapAdapter) getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//return zapcore.NewConsoleEncoder(encoderConfig)
	return zapcore.NewJSONEncoder(encoderConfig)
}

func (adapter *zapAdapter) setMaxSize(size int) {
	adapter.MaxSize = size
}

func (adapter *zapAdapter) setMaxBackups(n int) {
	adapter.MaxBackups = n
}

func (adapter *zapAdapter) setMaxAge(age int) {
	adapter.MaxAge = age
}

func (adapter *zapAdapter) setCompress(compress bool) {
	adapter.Compress = compress
}
func (adapter *zapAdapter) setCaller(caller bool) {
	adapter.Caller = caller
}

func (adapter *zapAdapter) setLogCategory(logCategory string) *zap.SugaredLogger {
	return adapter.sugar.Named(logCategory)
}

func (adapter *zapAdapter) WithFields(args ...interface{}) *zap.SugaredLogger {
	return adapter.sugar.With(args...)
}

func (adapter *zapAdapter) Debug(args ...interface{}) {
	adapter.sugar.Debug(args...)
}

func (adapter *zapAdapter) Info(args ...interface{}) {
	adapter.sugar.Info(args...)
}

func (adapter *zapAdapter) Warn(args ...interface{}) {
	adapter.sugar.Warn(args...)
}

func (adapter *zapAdapter) Error(args ...interface{}) {
	adapter.sugar.Error(args...)
}

func (adapter *zapAdapter) Panic(args ...interface{}) {
	adapter.sugar.Panic(args...)
}

func (adapter *zapAdapter) Fatal(args ...interface{}) {
	adapter.sugar.Fatal(args...)
}

func (adapter *zapAdapter) Debugf(template string, args ...interface{}) {
	adapter.sugar.Debugf(template, args...)
}

func (adapter *zapAdapter) Infof(template string, args ...interface{}) {
	adapter.sugar.Infof(template, args...)
}

func (adapter *zapAdapter) Warnf(template string, args ...interface{}) {
	adapter.sugar.Warnf(template, args...)
}

func (adapter *zapAdapter) Errorf(template string, args ...interface{}) {
	adapter.sugar.Errorf(template, args...)
}

func (adapter *zapAdapter) Panicf(template string, args ...interface{}) {
	adapter.sugar.Panicf(template, args...)
}

func (adapter *zapAdapter) Fatalf(template string, args ...interface{}) {
	adapter.sugar.Fatalf(template, args...)
}

func (adapter *zapAdapter) Debugw(msg string, keysAndValues ...interface{}) {
	adapter.sugar.Debugw(msg, keysAndValues...)
}

func (adapter *zapAdapter) Infow(msg string, keysAndValues ...interface{}) {
	adapter.sugar.Infow(msg, keysAndValues...)
}

func (adapter *zapAdapter) Warnw(msg string, keysAndValues ...interface{}) {
	adapter.sugar.Warnw(msg, keysAndValues...)
}

func (adapter *zapAdapter) Errorw(msg string, keysAndValues ...interface{}) {
	adapter.sugar.Errorw(msg, keysAndValues...)
}

func (adapter *zapAdapter) Panicw(msg string, keysAndValues ...interface{}) {
	adapter.sugar.Panicw(msg, keysAndValues...)
}

func (adapter *zapAdapter) Fatalw(msg string, keysAndValues ...interface{}) {
	adapter.sugar.Fatalw(msg, keysAndValues...)
}

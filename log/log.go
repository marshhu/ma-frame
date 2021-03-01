package log

const (
	DefaultLevel       = "info"
	DefaultPath        = "./logs"
	DefaultFileName    = "app.log"
	DefaultLogCategory = "app"
)

type LogSettings struct {
	Path        string // 日志文件路径
	FileName    string // 日志文件名
	Level       string // 日志输出的级别
	LogCategory string // 日志类别
	Caller      bool
	adapter     *zapAdapter
}

var logger *LogSettings

func Init(settings *LogSettings) error {
	logger = settings
	logger.adapter = NewZapAdapter(logger.Path, logger.FileName, logger.Level, logger.LogCategory)
	logger.adapter.setCaller(logger.Caller)
	logger.adapter.Build()
	return nil
}

// Sync flushes buffer, if any
func Sync() {
	if logger == nil {
		return
	}

	logger.adapter.logger.Sync()
}

// Debug 使用方法：log.Debug("test")
func Debug(args ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Debug(args...)
}

// Debugf 使用方法：log.Debugf("test:%s", err)
func Debugf(template string, args ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Debugf(template, args...)
}

// Debugw 使用方法：log.Debugw("test", "field1", "value1", "field2", "value2")
func Debugw(msg string, keysAndValues ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Debugw(msg, keysAndValues...)
}

func Info(args ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Info(args...)
}

func Infof(template string, args ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Infof(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Infow(msg, keysAndValues...)
}

func Warn(args ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Warnw(msg, keysAndValues...)
}

func Error(args ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Errorw(msg, keysAndValues...)
}

func Panic(args ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Panicw(msg, keysAndValues...)
}

func Fatal(args ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Fatalf(template, args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	if logger == nil {
		return
	}

	logger.adapter.Fatalw(msg, keysAndValues...)
}

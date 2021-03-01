package log

import "testing"

func TestDebug(t *testing.T) {
	Init(&LogSettings{
		Path:        DefaultPath,
		FileName:    DefaultFileName,
		Level:       "debug",
		LogCategory: DefaultLogCategory,
		Caller:      false,
	})

	Debug("test debug")
	Sync()
}

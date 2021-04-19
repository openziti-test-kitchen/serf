package testutil

import (
	"github.com/hashicorp/serf/logger"
	"io"
	"log"
	"strings"
	"testing"
)

func TestLogger(t testing.TB) logger.Logger {
	return logger.NewNamedLoggerWithFlags(&testWriter{t}, "test: ", log.LstdFlags)
}

func TestLoggerWithName(t testing.TB, name string) logger.Logger {
	return logger.NewNamedLoggerWithFlags(&testWriter{t}, "test["+name+"]: ", log.LstdFlags)
}

func TestWriter(t testing.TB) io.Writer {
	return &testWriter{t}
}

type testWriter struct {
	t testing.TB
}

func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.t.Helper()
	tw.t.Log(strings.TrimSpace(string(p)))
	return len(p), nil
}

package log

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setUpLoggerBufferForTesting() *bytes.Buffer {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	return &buf
}

func TestLoggerDebug(t *testing.T) {
	buffer := setUpLoggerBufferForTesting()
	defer log.SetOutput(os.Stdout)

	t.Run("Show args", func(t *testing.T) {
		t.Cleanup(func() { buffer.Reset() })

		SetLevel(DebugLevel)
		message := "message"
		expected := "[DEBUG] message test=123"

		Debug(message, "test", 123)

		assert.Contains(t, buffer.String(), expected)
	})

	t.Run("Show log when the level is Debug", func(t *testing.T) {
		t.Cleanup(func() { buffer.Reset() })

		SetLevel(DebugLevel)
		message := "message"
		expected := "[DEBUG] message"

		Debug(message)

		assert.Contains(t, buffer.String(), expected)
	})

	t.Run("Do not show log when the level is Info", func(t *testing.T) {
		t.Cleanup(func() { buffer.Reset() })

		SetLevel(InfoLevel)
		message := "message"

		Debug(message)

		assert.Empty(t, buffer.String())
	})
}

func TestLoggerInfo(t *testing.T) {
	buffer := setUpLoggerBufferForTesting()
	defer log.SetOutput(os.Stdout)

	t.Run("Show args", func(t *testing.T) {
		t.Cleanup(func() { buffer.Reset() })

		SetLevel(DebugLevel)
		message := "message"
		expected := "[INFO] message test=123"

		Info(message, "test", 123)

		assert.Contains(t, buffer.String(), expected)
	})

	t.Run("Show log when the level is Info", func(t *testing.T) {
		t.Cleanup(func() { buffer.Reset() })

		SetLevel(InfoLevel)
		message := "message"
		expected := "[INFO] message"

		Info(message)

		assert.Contains(t, buffer.String(), expected)
	})

	t.Run("Show log when the level is Debug", func(t *testing.T) {
		t.Cleanup(func() { buffer.Reset() })

		SetLevel(DebugLevel)
		message := "message"
		expected := "[INFO] message"

		Info(message)

		assert.Contains(t, buffer.String(), expected)
	})

	t.Run("Do not show log when the level is Warning", func(t *testing.T) {
		t.Cleanup(func() { buffer.Reset() })

		SetLevel(WarnLevel)
		message := "message"

		Info(message)

		assert.Empty(t, buffer.String())
	})
}

func TestLoggerWarn(t *testing.T) {
	buffer := setUpLoggerBufferForTesting()
	defer log.SetOutput(os.Stdout)

	t.Run("Show args", func(t *testing.T) {
		t.Cleanup(func() { buffer.Reset() })

		SetLevel(DebugLevel)
		message := "message"
		expected := "[WARN] message test=123"

		Warn(message, "test", 123)

		assert.Contains(t, buffer.String(), expected)
	})

	t.Run("Show log when the level is Warning", func(t *testing.T) {
		t.Cleanup(func() { buffer.Reset() })

		SetLevel(WarnLevel)
		message := "message"
		expected := "[WARN] message"

		Warn(message)

		assert.Contains(t, buffer.String(), expected)
	})

	t.Run("Show log when the level is Info", func(t *testing.T) {
		t.Cleanup(func() { buffer.Reset() })

		SetLevel(InfoLevel)
		message := "message"
		expected := "[WARN] message"

		Warn(message)

		assert.Contains(t, buffer.String(), expected)
	})

	t.Run("Do not show log when the level is Error", func(t *testing.T) {
		t.Cleanup(func() { buffer.Reset() })

		SetLevel(ErrorLevel)
		message := "message"

		Warn(message)

		assert.Empty(t, buffer.String())
	})
}

func TestLoggerError(t *testing.T) {
	buffer := setUpLoggerBufferForTesting()
	defer log.SetOutput(os.Stdout)

	t.Run("Show args", func(t *testing.T) {
		t.Cleanup(func() { buffer.Reset() })

		SetLevel(DebugLevel)
		message := "message"
		expected := "[ERROR] message test=123"

		Error(message, "test", 123)

		assert.Contains(t, buffer.String(), expected)
	})

	t.Run("Show log when the level is Error", func(t *testing.T) {
		t.Cleanup(func() { buffer.Reset() })

		SetLevel(ErrorLevel)
		message := "message"
		expected := "[ERROR] message"

		Error(message)

		assert.Contains(t, buffer.String(), expected)
	})

	t.Run("Show log when the level is Warning", func(t *testing.T) {
		t.Cleanup(func() { buffer.Reset() })

		SetLevel(WarnLevel)
		message := "message"
		expected := "[ERROR] message"

		Error(message)

		assert.Contains(t, buffer.String(), expected)
	})

	t.Run("Show log when the level is Info", func(t *testing.T) {
		t.Cleanup(func() { buffer.Reset() })

		SetLevel(InfoLevel)
		message := "message"
		expected := "[ERROR] message"

		Error(message)

		assert.Contains(t, buffer.String(), expected)
	})

	t.Run("Show log when the level is Debug", func(t *testing.T) {
		t.Cleanup(func() { buffer.Reset() })

		SetLevel(DebugLevel)
		message := "message"
		expected := "[ERROR] message"

		Error(message)

		assert.Contains(t, buffer.String(), expected)
	})
}

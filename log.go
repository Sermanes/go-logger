package log

import (
	"fmt"
	"log"
	"os"
)

const (
	// Levels
	FatalLevel uint8 = iota
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel

	separator = "="

	// Colors
	resetColor       = "\033[0m"
	redColor         = "\033[31m"
	greenColor       = "\033[32m"
	yellowColor      = "\033[33m"
	blueColor        = "\033[34m"
	magentaColor     = "\033[35m"
	cyanColor        = "\033[36m"
	whiteColor       = "\033[37m"
	blueBoldColor    = "\033[34;1m"
	magentaBoldColor = "\033[35;1m"
	redBoldColor     = "\033[31;1m"
	yellowBoldColor  = "\033[33;1m"
)

type Logger struct {
	level uint8
}

var (
	instance *Logger = &Logger{level: InfoLevel}
)

func Default() *Logger {
	return instance
}

func SetLevel(level uint8) {
	instance.level = level
}

func New(level uint8) *Logger {
	instance = &Logger{
		level: level,
	}

	return instance
}

func Debug(message string, args ...any) {
	if Default().level >= DebugLevel {
		log.Println(cyanColor, Default().log(fmt.Sprintf("[DEBUG] %s", message), args...), resetColor)
	}
}

func Info(message string, args ...any) {
	if Default().level >= InfoLevel {
		log.Println(blueColor, Default().log(fmt.Sprintf("[INFO] %s", message), args...), resetColor)
	}
}

func Warn(message string, args ...any) {
	if Default().level >= WarnLevel {
		log.Println(yellowColor, Default().log(fmt.Sprintf("[WARN] %s", message), args...), resetColor)
	}
}

func Error(message string, args ...any) {
	if Default().level >= ErrorLevel {
		log.Println(redColor, Default().log(fmt.Sprintf("[ERROR] %s", message), args...), resetColor)
	}
}

func Fatal(message string, args ...any) {
	defer os.Exit(1)
	log.Println(redBoldColor, Default().log(fmt.Sprintf("[FATAL] %s", message), args...), resetColor)
}

func (l *Logger) log(message string, args ...any) string {
	for index, arg := range args {
		if index%2 == 0 {
			message += fmt.Sprintf(" %s%s", arg, separator)
			continue
		}
		message += fmt.Sprintf("%v", arg)
	}

	return message
}

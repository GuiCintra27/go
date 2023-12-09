package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	err *log.Logger
}

func NewLogger (p string) *Logger {

	writer := io.Writer(os.Stdout)

	logger := log.New(writer, p, log.Ldate|log.Ltime)
	
	return &Logger{
		err: log.New(writer, "ERROR: ", logger.Flags()),
	}
}

//Create Non-Formated Log
func (l *Logger) Error(args ...interface{}) {
	l.err.Println(args...)
}

//Create Formated Enabled Log
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.err.Printf(format, args...)
}
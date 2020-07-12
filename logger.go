/*
#######
##         __
##        / /__  ___ ____ ____ ____
##       / / _ \/ _ `/ _ `/ -_) __/
##      /_/\___/\_, /\_, /\__/_/
##             /___//___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package logger

import (
	"bytes"
	"fmt"
	"os"
	"sync"
)

type (
	// Logger AFAIRE.
	Logger struct {
		bufPool *sync.Pool
		prefix  string
		level   Level
		fmt     Formatter
		out     Output
		clone   bool
	}
)

// New AFAIRE.
func New(prefix, level string, fmt Formatter, out Output) *Logger {
	return &Logger{
		bufPool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
		prefix: prefix,
		level:  GetLevelFromString(level),
		fmt:    fmt,
		out:    out,
	}
}

// Clone AFAIRE.
func (l *Logger) Clone(prefix string) *Logger {
	return &Logger{
		bufPool: l.bufPool,
		prefix:  prefix,
		level:   l.level,
		fmt:     l.fmt,
		out:     l.out,
		clone:   true,
	}
}

// Log AFAIRE.
func (l *Logger) Log(level Level, msg string, ctx ...interface{}) {
	if level < l.level {
		return
	}

	buf := l.bufPool.Get().(*bytes.Buffer)

	l.fmt.Format(buf, l.prefix, level, msg, l.out, ctx...)

	if err := l.out.Log(level, buf.Bytes()); err != nil {
		fmt.Fprintf( ///////////////////////////////////////////////////////////////////////////////////////////////////
			os.Stderr,
			"Output.Log() error >>> %s\n",
			err,
		)
	}

	buf.Reset()
	l.bufPool.Put(buf)
}

// Trace AFAIRE.
func (l *Logger) Trace(msg string, ctx ...interface{}) {
	l.Log(TraceLevel, msg, ctx...)
}

// Debug AFAIRE.
func (l *Logger) Debug(msg string, ctx ...interface{}) {
	l.Log(DebugLevel, msg, ctx...)
}

// Info AFAIRE.
func (l *Logger) Info(msg string, ctx ...interface{}) {
	l.Log(InfoLevel, msg, ctx...)
}

// Notice AFAIRE.
func (l *Logger) Notice(msg string, ctx ...interface{}) {
	l.Log(NoticeLevel, msg, ctx...)
}

// Warning AFAIRE.
func (l *Logger) Warning(msg string, ctx ...interface{}) {
	l.Log(WarningLevel, msg, ctx...)
}

// Error AFAIRE.
func (l *Logger) Error(msg string, ctx ...interface{}) {
	l.Log(ErrorLevel, msg, ctx...)
}

// Critical AFAIRE.
func (l *Logger) Critical(msg string, ctx ...interface{}) {
	l.Log(CriticalLevel, msg, ctx...)
}

// Close AFAIRE.
func (l *Logger) Close() {
	if !l.clone {
		if err := l.out.Close(); err != nil {
			fmt.Fprintf( ///////////////////////////////////////////////////////////////////////////////////////////////
				os.Stderr,
				"Output.Close() error >>> %s\n",
				err,
			)
		}
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/

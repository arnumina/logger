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
	"time"

	"github.com/arnumina/logfmt"
)

type (
	// TextFormatter AFAIRE.
	TextFormatter struct{}
)

// NewTextFormatter AFAIRE.
func NewTextFormatter() *TextFormatter {
	return &TextFormatter{}
}

// Format AFAIRE.
func (tf *TextFormatter) Format(
	buf *bytes.Buffer, prefix string, level Level, msg string, out Output, ctx ...interface{}) {
	if out.LogDateTime() {
		buf.WriteString(time.Now().Format("2006-01-02T15:04:05.000"))
	}

	if out.LogLevel() {
		switch level {
		case TraceLevel:
			buf.WriteString("{TRA} ")
		case DebugLevel:
			buf.WriteString("{DEB} ")
		case InfoLevel:
			buf.WriteString("{INF} ")
		case NoticeLevel:
			buf.WriteString("{NOT} ")
		case WarningLevel:
			buf.WriteString("{WAR} ")
		case ErrorLevel:
			buf.WriteString("{ERR} ")
		case CriticalLevel:
			buf.WriteString("{CRI} ")
		default:
			buf.WriteString("{???} ")
		}
	}

	buf.WriteString(prefix)
	buf.WriteRune(' ')
	buf.WriteString(msg)

	if len(ctx) != 0 {
		buf.WriteString("> ")
		logfmt.Encode(buf, ctx...)
	}

	if out.AddNewLine() {
		buf.WriteString("\n")
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/

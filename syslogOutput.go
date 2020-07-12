// +build !windows,!plan9

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

import "log/syslog"

type (
	// SyslogOutput AFAIRE.
	SyslogOutput struct {
		*syslog.Writer
		*OutputProperties
	}
)

func getFacilityFromString(facility string) syslog.Priority {
	switch facility {
	case "local0":
		return syslog.LOG_LOCAL0
	case "local1":
		return syslog.LOG_LOCAL1
	case "local2":
		return syslog.LOG_LOCAL2
	case "local3":
		return syslog.LOG_LOCAL3
	case "local4":
		return syslog.LOG_LOCAL4
	case "local5":
		return syslog.LOG_LOCAL5
	case "local6":
		return syslog.LOG_LOCAL6
	default:
		return syslog.LOG_LOCAL7
	}
}

// NewSyslogOutput AFAIRE.
func NewSyslogOutput(facility, app string) (Output, error) {
	writer, err := syslog.New(getFacilityFromString(facility), app)
	if err != nil {
		return nil, err
	}

	so := &SyslogOutput{
		Writer:           writer,
		OutputProperties: NewOutputProperties(false, false, false),
	}

	return so, nil
}

// Log AFAIRE.
func (so *SyslogOutput) Log(level Level, buf []byte) error {
	switch level {
	case TraceLevel:
		return so.Debug(string(buf))
	case DebugLevel:
		return so.Debug(string(buf))
	case InfoLevel:
		return so.Info(string(buf))
	case NoticeLevel:
		return so.Notice(string(buf))
	case WarningLevel:
		return so.Warning(string(buf))
	case ErrorLevel:
		return so.Err(string(buf))
	case CriticalLevel:
		return so.Crit(string(buf))
	default:
		return so.Debug(string(buf))
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/

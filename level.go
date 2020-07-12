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

type (
	// Level AFAIRE.
	Level int
)

const (
	// TraceLevel AFAIRE.
	TraceLevel Level = iota
	// DebugLevel AFAIRE.
	DebugLevel
	// InfoLevel AFAIRE.
	InfoLevel
	// NoticeLevel AFAIRE.
	NoticeLevel
	// WarningLevel AFAIRE.
	WarningLevel
	// ErrorLevel AFAIRE.
	ErrorLevel
	// CriticalLevel AFAIRE.
	CriticalLevel
)

// GetLevelFromString AFAIRE.
func GetLevelFromString(level string) Level {
	switch level {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "notice":
		return NoticeLevel
	case "warning":
		return WarningLevel
	case "error":
		return ErrorLevel
	case "critical":
		return CriticalLevel
	default:
		return TraceLevel
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/

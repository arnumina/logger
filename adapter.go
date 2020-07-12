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
	"log"
)

type (
	// LogAdapter AFAIRE.
	LogAdapter struct {
		level  Level
		logger *Logger
	}
)

// NewLogAdapter AFAIRE.
func (l *Logger) NewLogAdapter(level string) *LogAdapter {
	return &LogAdapter{
		level:  GetLevelFromString(level),
		logger: l,
	}
}

// Write AFAIRE.
func (la *LogAdapter) Write(p []byte) (int, error) {
	la.logger.Log(la.level, string(p))
	return len(p), nil
}

// NewLogger AFAIRE.
func (la *LogAdapter) NewLogger(prefix string, flag int) *log.Logger {
	return log.New(la, prefix, flag)
}

/*
######################################################################################################## @(°_°)@ #######
*/

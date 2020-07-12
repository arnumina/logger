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

import "bytes"

type (
	// Formatter AFAIRE
	Formatter interface {
		Format(buf *bytes.Buffer, prefix string, level Level, msg string, out Output, ctx ...interface{})
	}
)

/*
######################################################################################################## @(°_°)@ #######
*/

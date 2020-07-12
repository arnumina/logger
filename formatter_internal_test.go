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
	NopFormatter struct{}
)

func NewNopFormatter() *NopFormatter {
	return &NopFormatter{}
}

func (nf *NopFormatter) Format(_ *bytes.Buffer, _ string, _ Level, _ string, _ Output, _ ...interface{}) {
}

/*
######################################################################################################## @(°_°)@ #######
*/

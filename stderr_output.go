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

import "os"

type (
	// StderrOutput AFAIRE.
	StderrOutput struct {
		*WriterOutput
	}
)

// NewStderrOutput AFAIRE.
func NewStderrOutput() Output {
	return &StderrOutput{
		WriterOutput: NewWriterOutput(os.Stderr, NewOutputProperties(true, true, true)),
	}
}

// Close AFAIRE.
func (so *StderrOutput) Close() error {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/

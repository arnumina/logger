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
	// StdoutOutput AFAIRE.
	StdoutOutput struct {
		*WriterOutput
	}
)

// NewStdoutOutput AFAIRE.
func NewStdoutOutput() Output {
	return &StdoutOutput{
		WriterOutput: NewWriterOutput(os.Stdout, NewOutputProperties(true, true, true)),
	}
}

// Close AFAIRE.
func (so *StdoutOutput) Close() error {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/

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

import "io"

type (
	// Output AFAIRE.
	Output interface {
		LogDateTime() bool
		LogLevel() bool
		AddNewLine() bool
		Log(level Level, buf []byte) error
		Close() error
	}

	// OutputProperties AFAIRE.
	OutputProperties struct {
		dateTime bool
		level    bool
		newLine  bool
	}

	// WriterOutput AFAIRE.
	WriterOutput struct {
		io.Writer
		*OutputProperties
	}
)

// NewOutputProperties AFAIRE.
func NewOutputProperties(dateTime, level, newLine bool) *OutputProperties {
	return &OutputProperties{
		dateTime: dateTime,
		level:    level,
		newLine:  newLine,
	}
}

// LogDateTime AFAIRE.
func (op *OutputProperties) LogDateTime() bool {
	return op.dateTime
}

// LogLevel AFAIRE.
func (op *OutputProperties) LogLevel() bool {
	return op.level
}

// AddNewLine AFAIRE.
func (op *OutputProperties) AddNewLine() bool {
	return op.newLine
}

// NewWriterOutput AFAIRE.
func NewWriterOutput(iow io.Writer, op *OutputProperties) *WriterOutput {
	return &WriterOutput{
		Writer:           iow,
		OutputProperties: op,
	}
}

// Log AFAIRE.
func (wo *WriterOutput) Log(_ Level, buf []byte) error {
	_, err := wo.Write(buf)
	return err
}

/*
######################################################################################################## @(°_°)@ #######
*/

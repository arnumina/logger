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
	// FileOutput AFAIRE.
	FileOutput struct {
		*WriterOutput
		file *os.File
	}
)

// NewFileOutput AFAIRE.
func NewFileOutput(name string) (Output, error) {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	fo := &FileOutput{
		WriterOutput: NewWriterOutput(file, NewOutputProperties(true, true, true)),
		file:         file,
	}

	return fo, nil
}

// Close AFAIRE.
func (fo *FileOutput) Close() error {
	return fo.file.Close()
}

/*
######################################################################################################## @(°_°)@ #######
*/

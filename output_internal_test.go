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
	NopOutput struct {
		*OutputProperties
	}
)

func NewNopOutput() *NopOutput {
	return &NopOutput{
		OutputProperties: NewOutputProperties(false, false, false),
	}
}

func (no *NopOutput) Log(_ Level, _ []byte) error {
	return nil
}

func (no *NopOutput) Close() error {
	return nil
}

/*
type (
	BufOutput struct {
		*OutputProperties
		bytes.Buffer
	}
)

func NewBufOutput() *BufOutput {
	return &BufOutput{
		OutputProperties: NewOutputProperties(false, true, false),
	}
}

func (bo *BufOutput) Log(_ Level, buf []byte) error {
	_, err := bo.Write(buf)
	return err
}

func (bo *BufOutput) Close() error {
	return nil
}
*/

/*
######################################################################################################## @(°_°)@ #######
*/

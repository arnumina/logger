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

package logger_test

import (
	"testing"

	"github.com/arnumina/logger"
)

var (
	_fakeMessage = "Test logging, but use a somewhat realistic message length."
)

func newLogger(fmt logger.Formatter) *logger.Logger {
	return logger.New("benchmark", "info", fmt, logger.NewNopOutput())
}

func BenchmarkNop(b *testing.B) {
	logger := newLogger(logger.NewNopFormatter())

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		logger.Info(_fakeMessage)
	}
}

func BenchmarkDisabled(b *testing.B) {
	logger := newLogger(logger.NewNopFormatter())

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		logger.Debug(_fakeMessage) // debugLevel < infoLevel
	}
}

func BenchmarkWithoutContext(b *testing.B) {
	logger := newLogger(logger.NewTextFormatter())

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		logger.Info(_fakeMessage)
	}
}

func BenchmarkWithContext(b *testing.B) {
	logger := newLogger(logger.NewTextFormatter())

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		logger.Info(_fakeMessage, "day", 12, "month", "April", "year", 2020)
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/

package g_met

import (
	"github.com/cihub/seelog"
)

type SeeLogWriter struct {
	Logger seelog.LoggerInterface
}

func (writer *SeeLogWriter) Write(msg string) {
	writer.Logger.Info(msg)
}

func (writer *SeeLogWriter) Flush() {
	writer.Logger.Flush()
}

func (writer *SeeLogWriter) Close() error {
	return writer.Close()
}

func CreateMetWriterBySeeLog(configFile string) (MetWriter, error) {
	var err error
	writer := new(SeeLogWriter)
	writer.Logger, err = seelog.LoggerFromConfigAsFile(configFile)
	if err != nil {
		return nil, err
	}

	return writer, nil
}

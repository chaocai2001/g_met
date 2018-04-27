package g_met

type GMetInstance struct {
	metWriter    MetWriter
	metFormatter MetFormatter
}

func CreateGMetInstance(metWriter MetWriter, metFormatter MetFormatter) GMet {
	ins := GMetInstance{metWriter, metFormatter}
	return &ins
}

func (gmet *GMetInstance) Send(metrics ...MetricItem) error {
	if formatted, err := gmet.metFormatter.Format(metrics); err != nil {
		return err
	} else {
		gmet.metWriter.Write(formatted)
	}
	return nil
}

func (gmet *GMetInstance) Flush() {
	gmet.metWriter.Flush()
}

func (gmet *GMetInstance) Close() error {
	return gmet.metWriter.Close()
}

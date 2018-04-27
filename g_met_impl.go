package g_met

import (
	"net"
)

const (
	HOST_ADDR     = "HostAddr"
	MISSING_VALUE = "N/A"
)

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

//Get the local IP adress
func IpAddress() (MetricItem, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return MetricItem{HOST_ADDR, MISSING_VALUE}, err
	}
	for _, address := range addrs {
		// Check if it is ip circle
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return MetricItem{HOST_ADDR, ipnet.IP.String()}, nil
			}

		}
	}
	return MetricItem{HOST_ADDR, MISSING_VALUE}, err
}

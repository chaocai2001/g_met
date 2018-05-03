//GMet is golang client of XMet API.
//For more, see also: https://github.com/chaocai2001/g_met
//Created on 2018.5
package g_met

import (
	"net"
)

const (
	HOST_ADDR     = "HostAddr"
	MISSING_VALUE = "N/A"
)

type GMetInstance struct {
	metWriter    MetWriter    //metrics data writer
	metFormatter MetFormatter //metrics formatter
}

func CreateGMetInstance(metWriter MetWriter, metFormatter MetFormatter) GMet {
	ins := GMetInstance{metWriter, metFormatter}
	return &ins
}

//Create GMet Instance with default settings.
//(with seelog writer and ltr format
func CreateGMetInstanceByDefault(metricsFile string) GMet {
	//create a metric writer
	writer, err := CreateMetWriterBySeeLog(metricsFile)
	if err != nil {
		panic(err)
	}
	//create GMet instance by given the writer and the formatter
	gmet := CreateGMetInstance(writer, &LtrFormatter{})
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

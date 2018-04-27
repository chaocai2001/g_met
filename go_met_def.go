package g_met

type MetricItem struct {
	Key   string
	Value string
}

type GMet interface {
	Send(metrics ...MetricItem) error
	Flush()
	Close() error
}

type MetWriter interface {
	Write(msg string)
	Flush()
	Close() error
}

type MetFormatter interface {
	Format(metrics []MetricItem) (string, error)
}

func Metric(key string, value string) MetricItem {
	return MetricItem{key, value}
}

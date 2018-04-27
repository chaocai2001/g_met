package g_met

import (
	"bytes"
	"strconv"
	"strings"
	"time"
)

const SEC_SPLITTER string = "\t"
const FIELD_SPLITTER string = ":"
const TIMESTAMP_KEY string = "timestamp"

type LtrFormatter struct{}

func preprocess(value string) string {
	p1 := strings.Replace(value, SEC_SPLITTER, "_", -1)
	p2 := strings.Replace(p1, FIELD_SPLITTER, "-", -1)
	return p2
}

func (formatter *LtrFormatter) Format(metrics []MetricItem) (string, error) {
	buf := bytes.NewBufferString("")
	buf.WriteString(TIMESTAMP_KEY)
	buf.WriteString(FIELD_SPLITTER)
	buf.WriteString(strconv.FormatInt(time.Now().Unix(), 10))
	buf.WriteString(SEC_SPLITTER)
	for _, metric := range metrics {
		buf.WriteString(preprocess(metric.Key))
		buf.WriteString(FIELD_SPLITTER)
		buf.WriteString(preprocess(metric.Value))
		buf.WriteString(SEC_SPLITTER)
	}
	return buf.String(), nil
}

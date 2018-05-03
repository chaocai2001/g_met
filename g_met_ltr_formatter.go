//The format is [Metric Name][Field Splitter][Metric Value][Section Splitter][Metric Name][Field Splitter][Metric Value]
//Created on 2018.5
package g_met

import (
	"bytes"
	"strconv"
	"strings"
	"time"
)

const (
	SEC_SPLITTER               string = "\t"
	FIELD_SPLITTER             string = ":"
	TIMESTAMP_KEY              string = "timestamp"
	SEC_SPLITTER_REPLACEMENT   string = "_"
	FIELD_SPLITTER_REPLACEMENT string = "-"
)

type LtrFormatter struct{}

func replaceSplitterCharsInValue(value string) string {
	p1 := strings.Replace(value, SEC_SPLITTER, SEC_SPLITTER_REPLACEMENT, -1)
	p2 := strings.Replace(p1, FIELD_SPLITTER, FIELD_SPLITTER_REPLACEMENT, -1)
	return p2
}

func (formatter *LtrFormatter) Format(metrics []MetricItem) (string, error) {
	buf := bytes.NewBufferString("")
	buf.WriteString(TIMESTAMP_KEY)
	buf.WriteString(FIELD_SPLITTER)
	buf.WriteString(strconv.FormatInt(time.Now().Unix(), 10))
	buf.WriteString(SEC_SPLITTER)
	for _, metric := range metrics {
		buf.WriteString(replaceSplitterCharsInValue(metric.Key))
		buf.WriteString(FIELD_SPLITTER)
		buf.WriteString(replaceSplitterCharsInValue(metric.Value))
		buf.WriteString(SEC_SPLITTER)
	}
	return buf.String(), nil
}

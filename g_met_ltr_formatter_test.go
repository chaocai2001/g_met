package g_met

import (
	"strings"
	"testing"
)

func TestFormatMetricItemWithNoSplitterInValue(t *testing.T) {
	formatter := LtrFormatter{}
	metricItems := []MetricItem{
		MetricItem{"K1", "V1"},
		MetricItem{"K2", "V2"},
	}
	formatted, err := formatter.Format(metricItems)
	if err != nil {
		t.Errorf("failed to format %v", err)
	}
	expected := "K1" + FIELD_SPLITTER + "V1" +
		SEC_SPLITTER + "K2" + FIELD_SPLITTER + "V2"
	if !strings.Contains(formatted, expected) {
		t.Errorf("The formatted is %s, but the items is expected as %s\n",
			formatted, expected)
	}

	if !strings.Contains(formatted, TIMESTAMP_KEY) {
		t.Error("Timestamp is missing\n")
	}
}

func TestFormatMetricItemWithSplitterInValue(t *testing.T) {
	formatter := LtrFormatter{}
	metricItems := []MetricItem{
		MetricItem{"K1\t", "V1:A"},
		MetricItem{"K2:G", "V2\t"},
	}
	formatted, err := formatter.Format(metricItems)
	if err != nil {
		t.Errorf("failed to format %v", err)
	}
	expected := "K1_" + FIELD_SPLITTER + "V1-A" +
		SEC_SPLITTER + "K2-G" + FIELD_SPLITTER + "V2"
	if !strings.Contains(formatted, expected) {
		t.Errorf("The formatted is %s, but the items is expected as %s\n",
			formatted, expected)
	}

	if !strings.Contains(formatted, TIMESTAMP_KEY) {
		t.Error("Timestamp is missing\n")
	}
}

package metric_model

import (
	"github.com/conductor-sdk/conductor-go/metrics/metric_model/metric_documentation"
	"github.com/conductor-sdk/conductor-go/metrics/metric_model/metric_label"
	"github.com/conductor-sdk/conductor-go/metrics/metric_model/metric_name"
)

type MetricDetails struct {
	Name        string
	Description string
	Labels      []string
}

func NewMetricDetails(
	name metric_name.MetricName,
	description metric_documentation.MetricDocumentation,
	labels []metric_label.MetricLabel,
) *MetricDetails {
	metricDetails := new(MetricDetails)
	metricDetails.Name = string(name)
	metricDetails.Description = string(description)
	for _, label := range labels {
		metricDetails.Labels = append(metricDetails.Labels, string(label))
	}
	return metricDetails
}

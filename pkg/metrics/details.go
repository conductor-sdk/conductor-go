package metrics

type MetricDetails struct {
	Name        string
	Description string
	Labels      []string
}

func NewMetricDetails(
	name MetricName,
	description MetricDocumentation,
	labels []MetricLabel,
) *MetricDetails {
	metricDetails := new(MetricDetails)
	metricDetails.Name = string(name)
	metricDetails.Description = string(description)
	for _, label := range labels {
		metricDetails.Labels = append(metricDetails.Labels, string(label))
	}
	return metricDetails
}

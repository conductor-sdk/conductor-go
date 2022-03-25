package metrics

type MetricDetails struct {
	Name        string
	Description string
	Labels      []string
}

func NewMetricDetails(name string, description string, labels []string) *MetricDetails {
	metricDetails := new(MetricDetails)
	metricDetails.Name = name
	metricDetails.Description = description
	metricDetails.Labels = labels
	return metricDetails
}

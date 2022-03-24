package metrics

import (
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func ProvideMetrics(pattern string, port int) {
	http.Handle(pattern, promhttp.Handler())
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

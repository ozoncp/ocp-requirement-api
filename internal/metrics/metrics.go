package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	createCounter prometheus.Counter
	updateCounter prometheus.Counter
	removeCounter prometheus.Counter
)

func InitMetrics() {
	createCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocp_requirement_api_create",
	})

	updateCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocp_requirement_api_update",
	})

	removeCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ocp_requirement_api_remove",
	})
}

func CreateCounterInc() {
	createCounter.Inc()
}

func UpdateCounterInc() {
	updateCounter.Inc()
}

func RemoveCounterInc() {
	removeCounter.Inc()
}

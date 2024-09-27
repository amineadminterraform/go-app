package prom_metrics

import "github.com/prometheus/client_golang/prometheus"
type Device struct {
	ID int `json:"id"`
	Mac string `json:"mac"`
	Firemware string `json:"firemware"`
}
type metrics struct {
	Devices prometheus.Gauge
	Info *prometheus.GaugeVec
}
func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		Devices: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "app",
			Name:      "devices",
			Help:      "Number of currently connected devices",
		}),
		Info: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "app",
			Name: "info",
			Help: "Informations about the app",
		},
		[]string{"version"}),
	
	}



	reg.MustRegister(m.Devices,m.Info)
	return m
		
}

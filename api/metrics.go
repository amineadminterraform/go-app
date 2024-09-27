package api

import (
	"encoding/json"
	"net/http"

	prom_metrics "github.com/amineadminterraform/go-app/pkg/metrics"
)

// var devices []metrics.Device = []metrics.Device{}
var version string
var devices []prom_metrics.Device = []prom_metrics.Device{
	{ID: 1, Mac: "00:00:10:00:00:01", Firemware: "1.0.0"},
	{ID: 2, Mac: "01:00:00:00:00:01", Firemware: "2.0.0"},
	{ID: 3, Mac: "10:00:00:00:00:01", Firemware: "3.0.0"},
}

func GetDevices(w http.ResponseWriter, r *http.Request){
	b, err := json.Marshal(devices)
	if err != nil { 
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)	
	w.Write(b)
} 

func CreateDevice(w http.ResponseWriter, r *http.Request){
	var dv prom_metrics.Device
	err := json.NewDecoder(r.Body).Decode(&dv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	devices = append(devices, dv)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Device created"))
}


// func SetupMetrics() http.Handler{
	
// 	version := "2.16.5"
// 	req := prometheus.NewRegistry()
// 	m := prom_metrics.NewMetrics(req)
// 	m.Info.With(prometheus.Labels{"version": version}).Set(1)
// 	m.Devices.Set(float64(len(devices)))
// 	promHandler := promhttp.HandlerFor(req, promhttp.HandlerOpts{} )
// 	return promHandler
// }


package main

import (
	"encoding/json"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func sum(input []float64) float64 {
	sum := 0.0
	for i := range input {
		sum += input[i]
	}
	return sum
}

// JsonData struct
type JsonData struct {
	UUID      string    `json:"id"`
	Numbers   []float64 `json:"numbers"`
	Operation string    `json:"operation"`
	Status    string    `json:"status"`
}

type Message struct {
	UUID      string    `json:"id,omitempty"`
	Numbers   []float64 `json:"numbers,omitempty"`
	Result    float64   `json:"result,omitempty"`
	Status    string    `json:"status,omitempty"`
	Operation string    `json:"operation,omitempty"`
}

func process(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data JsonData
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	var msg Message
	msg.UUID = data.UUID
	msg.Numbers = data.Numbers
	msg.Operation = data.Operation
	msg.Status = "done"
	msg.Result = sum(data.Numbers)
	output, err := json.Marshal(msg)
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}
	rw.Header().Set("content-type", "application/json")
	rw.Write(output)
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/operation", process)
	panic(http.ListenAndServe(":8080", nil))
}

package hmetrics2graphite

import (
	"fmt"
	"github.com/marpaia/graphite-golang"
	"math"
)

func Exporter(host string, port int, prefix string) (func(map[string]float64), error) {
	gr, err := graphite.NewGraphite(host, port)
	if err != nil {
		return nil, err
	}
	err = gr.Connect()
	if err != nil {
		return nil, err
	}
	return func(newData map[string]float64) {
		for k, v := range newData {
			if !math.IsNaN(v) && !math.IsInf(v, 0) {
				gr.SimpleSend(prefix+"."+k, fmt.Sprintf("%v", v))
			}
		}
	}, nil
}

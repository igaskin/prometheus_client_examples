// Learn more or give us feedback
// Copyright 2015 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// A minimal example of how to include Prometheus instrumentation.
package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var addr = flag.String("listen-address", "0.0.0.0:8080", "The address to listen on for HTTP requests.")

var (
	rpcExampleState = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "rpc_state_count",
			Help: "RPC state counter",
		},
		[]string{"status", "service"},
	)
)

func init() {
	prometheus.MustRegister(rpcExampleState)
}

func main() {
	flag.Parse()

	go func() {
		for {
			time.Sleep(time.Second * 1)
			rpcExampleState.WithLabelValues("success", "factory-handler").Inc()
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}

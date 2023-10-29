package prom

// Copyright 2023 Ahmed Abdelkafi
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

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/razerban/linky-exporter/pkg/core"
	log "github.com/sirupsen/logrus"
)

// LinkyExporter object to run exporter server and expose metrics
type LinkyExporter struct {
	Address string
	Port    int
}

// Run method to run http exporter server
func (exporter *LinkyExporter) Run(connector core.LinkyConnector) {
	log.Info(fmt.Sprintf("Beginning to serve on port :%d", exporter.Port))

	prometheus.MustRegister(NewLinkyCollector(connector))
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", exporter.Address, exporter.Port), nil))
}

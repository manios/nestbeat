package status

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/metricbeat/mb"
)

// init registers the MetricSet with the central registry.
// The New method will be called after the setup of the module and before starting to fetch data
func init() {
	if err := mb.Registry.AddMetricSet("thermostat", "status", New); err != nil {
		panic(err)
	}
}

// MetricSet type defines all fields of the MetricSet
// As a minimum it must inherit the mb.BaseMetricSet fields, but can be extended with
// additional entries. These variables can be used to persist data or configuration between
// multiple fetch calls.
type MetricSet struct {
	mb.BaseMetricSet
	// OAuth access token which has to be used in order to authenticate against Nest API
	accessToken string
	// Base url of Nest API
	apiHost string
	// Nest Thermostat unique identifier
	deviceID string
}

// New create a new instance of the MetricSet
// Part of new is also setting up the configuration by processing additional
// configuration entries if needed.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {

	// Unpack additional configuration options
	config := struct {
		accessToken string `config:"accessToken"`
		apiHost     string `config:"apiHost"`
		deviceID    string `config:"deviceID"`
	}{
		accessToken: "",
		apiHost:     "https://developer-api.nest.com",
		deviceID:    "",
	}

	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	return &MetricSet{
		BaseMetricSet: base,
		accessToken:   config.accessToken,
		apiHost:       config.apiHost,
		deviceID:      config.deviceID,
	}, nil
}

// Fetch methods implements the data gathering and data conversion to the right format
// It returns the event which is then forward to the output. In case of an error, a
// descriptive error must be returned.
func (m *MetricSet) Fetch() (common.MapStr, error) {

	event := common.MapStr{
	// "counter": m.counter,
	}
	// m.counter++

	return event, nil
}

package status

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/metricbeat/mb"
	"github.com/manios/nestbeat/module/thermostat/status/gonest"
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
	deviceID    string
	nestAPIConf *gonest.Configuration
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

	// Create a new Nest API configuration according to beat configuration
	nestAPIConf := gonest.NewConfiguration()
	nestAPIConf.APIKey["Authorization"] = "Bearer " + config.accessToken
	nestAPIConf.BasePath = config.apiHost

	return &MetricSet{
		BaseMetricSet: base,
		accessToken:   config.accessToken,
		apiHost:       config.apiHost,
		deviceID:      config.deviceID,
		nestAPIConf:   nestAPIConf,
	}, nil
}

// Fetch methods implements the data gathering and data conversion to the right format
// It returns the event which is then forward to the output. In case of an error, a
// descriptive error must be returned.
func (m *MetricSet) Fetch() (common.MapStr, error) {

	var apos = gonest.NewThermostatApi()
	apos.Configuration = m.nestAPIConf

	// Retrieve stats for Nest thermostat
	tstat, _, _ := apos.DevicesThermostatsThermostatUidGet(m.deviceID)

	// transform stats to event
	event := common.MapStr{
		"humidity":                    tstat.Humidity,
		"locale":                      tstat.Locale,
		"temperature_scale":           tstat.TemperatureScale,
		"is_using_emergency_heat":     tstat.IsUsingEmergencyHeat,
		"has_fan":                     tstat.HasFan,
		"software_version":            tstat.SoftwareVersion,
		"has_leaf":                    tstat.HasLeaf,
		"where_id":                    tstat.WhereId,
		"device_id":                   tstat.DeviceId,
		"name":                        tstat.Name,
		"can_heat":                    tstat.CanHeat,
		"can_cool":                    tstat.CanCool,
		"target_temperature_c":        tstat.TargetTemperatureC,
		"target_temperature_f":        tstat.TargetTemperatureF,
		"target_temperature_high_c":   tstat.TargetTemperatureHighC,
		"target_temperature_high_f":   tstat.TargetTemperatureHighF,
		"target_temperature_low_c":    tstat.TargetTemperatureLowC,
		"target_temperature_low_f":    tstat.TargetTemperatureLowF,
		"ambient_temperature_c":       tstat.AmbientTemperatureC,
		"ambient_temperature_f":       tstat.AmbientTemperatureF,
		"away_temperature_high_c":     tstat.AwayTemperatureHighC,
		"away_temperature_high_f":     tstat.AwayTemperatureHighF,
		"away_temperature_low_c":      tstat.AwayTemperatureLowC,
		"away_temperature_low_f":      tstat.AwayTemperatureLowF,
		"eco_temperature_high_c":      tstat.EcoTemperatureHighC,
		"eco_temperature_high_f":      tstat.EcoTemperatureHighF,
		"eco_temperature_low_c":       tstat.EcoTemperatureLowC,
		"eco_temperature_low_f":       tstat.EcoTemperatureLowF,
		"is_locked":                   tstat.IsLocked,
		"locked_temp_min_c":           tstat.LockedTempMinC,
		"locked_temp_min_f":           tstat.LockedTempMinF,
		"locked_temp_max_c":           tstat.LockedTempMaxC,
		"locked_temp_max_f":           tstat.LockedTempMaxF,
		"sunlight_correction_active":  tstat.SunlightCorrectionActive,
		"sunlight_correction_enabled": tstat.SunlightCorrectionEnabled,
		"structure_id":                tstat.StructureId,
		"fan_timer_active":            tstat.FanTimerActive,
		"fan_timer_timeout":           tstat.FanTimerTimeout,
		"fan_timer_duration":          tstat.FanTimerDuration,
		"previous_hvac_mode":          tstat.PreviousHvacMode,
		"hvac_mode":                   tstat.HvacMode,
		"time_to_target":              tstat.TimeToTarget,
		"time_to_target_training":     tstat.TimeToTargetTraining,
		"where_name":                  tstat.WhereName,
		"label":                       tstat.Label,
		"name_long":                   tstat.NameLong,
		"is_online":                   tstat.IsOnline,
		"hvac_state":                  tstat.HvacState,
	}

	return event, nil
}

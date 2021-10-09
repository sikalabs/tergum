package telemetry

import (
	"encoding/json"
	"os"
	"runtime"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/sikalabs/tergum/version"
)

const DEFAULT_TELEMETRY_ORIGIN = "https://tergum-telemetry-api.sikalabs.com"

type TelemetryConfig struct {
	Origin          string `yaml:"Origin"`
	Disable         bool   `yaml:"Disable"`
	Name            string `yaml:"Name"`
	CollectHostData bool   `yaml:"CollectHostData"`
}

type HostData struct {
	Hostname string
	GOOS     string
	GOARCH   string
	HostInfo host.InfoStat
}

type Telemetry struct {
	Config   TelemetryConfig
	Enabled  bool
	Client   *resty.Client
	HostData *HostData
}

func NewTelemetry(tc *TelemetryConfig, disabled bool, extraName string) Telemetry {
	if tc == nil {
		tc = &TelemetryConfig{}
	}

	if tc.Disable || disabled {
		log.Info().
			Msg("Telemetry disabled.")
		return Telemetry{
			Enabled: false,
		}
	}

	if tc.Origin == "" {
		tc.Origin = DEFAULT_TELEMETRY_ORIGIN
	}

	if extraName != "" {
		tc.Name = extraName
	}

	client := resty.New()
	client.SetTimeout(5 * time.Second)

	log.Info().
		Str("origin", tc.Origin).
		Str("telemetry_name", tc.Name).
		Str("version", version.Version).
		Msg("Telemetry backend initialized.")

	var hostData *HostData
	if tc.CollectHostData {
		hostname, _ := os.Hostname()
		hostInfo, _ := host.Info()
		hostData = &HostData{
			Hostname: hostname,
			GOOS:     runtime.GOOS,
			GOARCH:   runtime.GOARCH,
			HostInfo: *hostInfo,
		}
	}

	return Telemetry{
		Enabled:  true,
		Config:   *tc,
		Client:   client,
		HostData: hostData,
	}
}

func (t Telemetry) GetHostname() string {
	if t.HostData != nil {
		return t.HostData.Hostname
	}
	return ""
}

func (t *Telemetry) SendEvent(name, data string) {
	if !t.Enabled {
		log.Info().
			Msg("Telemetry skip.")
		return
	}
	_, err := t.Client.R().
		SetBody(map[string]interface{}{
			"version":        version.Version,
			"telemetry_name": t.Config.Name,
			"hostname":       t.GetHostname(),
			"event_name":     name,
			"data":           data,
		}).
		Post(t.Config.Origin + "/api/v1/event")
	if err == nil {
		log.Info().
			Msg("Telemetry successfully sent.")
	} else {
		log.Warn().
			Msg("Telemetry failed.")
	}
}

func (t *Telemetry) SendEventInit() {
	out, _ := json.Marshal(t.HostData)
	t.SendEvent("init/v2", string(out))
}

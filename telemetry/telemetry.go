package telemetry

import (
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"github.com/sikalabs/tergum/version"
)

const DEFAULT_TELEMETRY_ORIGIN = "https://tergum-telemetry-api.sikalabs.com"

type TelemetryConfig struct {
	Origin  string `yaml:"Origin"`
	Disable bool   `yaml:"Disable"`
	Name    string `yaml:"Name"`
}

type Telemetry struct {
	Config  TelemetryConfig
	Enabled bool
	Client  *resty.Client
}

func NewTelemetry(tc *TelemetryConfig, disabled bool, extraName string) Telemetry {
	if tc == nil {
		tc = &TelemetryConfig{
			Origin: DEFAULT_TELEMETRY_ORIGIN,
		}
	}

	if tc.Disable || disabled {
		log.Info().
			Msg("Telemetry disabled.")
		return Telemetry{
			Enabled: false,
		}
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

	return Telemetry{
		Enabled: true,
		Config:  *tc,
		Client:  client,
	}
}

func (t *Telemetry) SendTelemetry() {
	if !t.Enabled {
		log.Info().
			Msg("Telemetry skip.")
		return
	}
	_, err := t.Client.R().
		SetBody(map[string]interface{}{
			"version":        version.Version,
			"telemetry_name": t.Config.Name,
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

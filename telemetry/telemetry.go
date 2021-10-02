package telemetry

import (
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"github.com/sikalabs/tergum/version"
)

const DEFAULT_TELEMETRY_ORIGIN = "https://tergum-telemetry-api.sikalabs.com"

type Telemetry struct {
	Enabled bool
	Origin  string
	Name    string
	Client  *resty.Client
}

func NewTelemetry(enabled bool, origin string, name string) Telemetry {
	if !enabled {
		log.Info().
			Msg("Telemetry disabled.")
		return Telemetry{
			Enabled: enabled,
		}
	}

	if origin == "" {
		origin = DEFAULT_TELEMETRY_ORIGIN
	}
	client := resty.New()
	client.SetTimeout(5 * time.Second)
	log.Info().
		Str("origin", origin).
		Str("telemetry_name", name).
		Msg("Telemetry backend initialized.")
	return Telemetry{
		Enabled: enabled,
		Origin:  origin,
		Name:    name,
		Client:  client,
	}
}

func (tc *Telemetry) SendTelemetry() {
	if !tc.Enabled {
		log.Info().
			Msg("Telemetry skip.")
		return
	}
	_, err := tc.Client.R().
		SetBody(map[string]interface{}{
			"version":        version.Version,
			"telemetry_name": tc.Name,
		}).
		Post(tc.Origin + "/api/v1/event")
	if err == nil {
		log.Info().
			Msg("Telemetry successfully sent.")
	} else {
		log.Warn().
			Msg("Telemetry failed.")
	}
}

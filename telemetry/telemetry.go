package telemetry

import (
	"encoding/json"
	"os"
	"runtime"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"github.com/sikalabs/tergum/backup_log"
	"github.com/sikalabs/tergum/version"
)

const DEFAULT_TELEMETRY_ORIGIN = "https://tergum-telemetry-api.sikalabs.com"

type TelemetryConfig struct {
	Origin           string `yaml:"Origin"`
	Disable          bool   `yaml:"Disable"`
	Name             string `yaml:"Name"`
	CollectHostData  bool   `yaml:"CollectHostData"`
	CollectBackupLog bool   `yaml:"CollectBackupLog"`
}

type HostData struct {
	Hostname string
	GOOS     string
	GOARCH   string
}

type Telemetry struct {
	Config     TelemetryConfig
	Enabled    bool
	Client     *resty.Client
	HostData   *HostData
	CloudEmail string
}

func NewTelemetry(
	tc *TelemetryConfig,
	disabled bool,
	extraName string,
	cloudEmail string,
) Telemetry {
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
		Str("cloud_email", cloudEmail).
		Msg("Telemetry backend initialized.")

	var hostData *HostData
	if tc.CollectHostData {
		hostname, _ := os.Hostname()
		hostData = &HostData{
			Hostname: hostname,
			GOOS:     runtime.GOOS,
			GOARCH:   runtime.GOARCH,
		}
	}

	return Telemetry{
		Enabled:    true,
		Config:     *tc,
		Client:     client,
		HostData:   hostData,
		CloudEmail: cloudEmail,
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
			Str("event_name", name).
			Msg("Telemetry skip.")
		return
	}
	_, err := t.Client.R().
		SetBody(map[string]interface{}{
			"version":        version.Version,
			"telemetry_name": t.Config.Name,
			"cloud_email":    t.CloudEmail,
			"hostname":       t.GetHostname(),
			"event_name":     name,
			"data":           data,
		}).
		Post(t.Config.Origin + "/api/v1/event")
	if err == nil {
		log.Debug().
			Str("event_name", name).
			Msg("Telemetry successfully sent.")
	} else {
		log.Warn().
			Str("event_name", name).
			Msg("Telemetry failed.")
	}
}

func (t *Telemetry) SendEventInit() {
	out, _ := json.Marshal(t.HostData)
	t.SendEvent("init/v2", string(out))
}

func (t *Telemetry) SendEventBackupLog(bl backup_log.BackupLog) {
	out, _ := json.Marshal(bl)
	t.SendEvent("backup-log-dump/v1", string(out))
}

func (t *Telemetry) SendEventInitExtra(doBackupImplementation string) {
	out, _ := json.Marshal(map[string]string{
		"do_backup_implementation": doBackupImplementation,
	})
	t.SendEvent("init-extra/v1", string(out))
}

package cloud

const DEFAULT_CLOUD_ORIGIN = "https://tergum-cloud-api.sikalabs.com"

type CloudConfig struct {
	Email string `yaml:"Email" json:"Email,omitempty"`
}

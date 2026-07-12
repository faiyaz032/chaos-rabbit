package config

type Config struct {
	HTTP  HTTPConfig
	Chaos []ChaosConfig
}

type HTTPConfig struct {
	Listen string
	Target string
}

type ChaosConfig struct {
	Type   string         `yaml:"type"`
	Config map[string]any `yaml:"config"`
}

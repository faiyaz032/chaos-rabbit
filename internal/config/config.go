package config

type Config struct {
	HTTP HTTPConfig
}

type HTTPConfig struct {
	Listen string
	Target string
}

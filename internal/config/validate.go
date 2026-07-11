package config

import "fmt"

func (c Config) Validate() error {

	if c.HTTP.Listen == "" {
		return fmt.Errorf("http listen address required")
	}

	if c.HTTP.Target == "" {
		return fmt.Errorf("http target required")
	}

	return nil
}

package biztalkserverreceiver

import (
	"fmt"
	"time"

	"go.opentelemetry.io/collector/scraper/scraperhelper"
)

type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	Endpoint                       string `mapstructure:"endpoint"`
	Interval                       string `mapstructure:"interval"`
	Username                       string `mapstructure:"username"` // TODO: Refactor into auth sub struct?
	Password                       string `mapstructure:"password"`
}

// Validate checks if the receiver configuration is valid
func (cfg *Config) Validate() error {
	interval, _ := time.ParseDuration(cfg.Interval)
	if interval.Minutes() < 1 {
		return fmt.Errorf("when defined, the interval has to be set to at least 1 minute (1m)")
	}
	if cfg.Endpoint == "" {
		return fmt.Errorf("Endpoint must not be empty")
	}
	return nil
}

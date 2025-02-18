package biztalkserverreceiver

import (
	"fmt"
	"github.com/Integrio/biztalkserverreceiver/biztalkserverreceiver/internal/metadata"
	"go.opentelemetry.io/collector/scraper/scraperhelper"
)

type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	metadata.MetricsBuilderConfig  `mapstructure:",squash"`
	Endpoint                       string `mapstructure:"endpoint"`
	Auth                           string `mapstructure:"auth"`
	Username                       string `mapstructure:"username"`
	Password                       string `mapstructure:"password"`
}

// Validate checks if the receiver configuration is valid
func (cfg *Config) Validate() error {
	if cfg.Endpoint == "" {
		return fmt.Errorf("endpoint must not be empty")
	}
	return nil
}

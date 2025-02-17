package biztalkserverreceiver

import (
	"context"
	"github.com/Integrio/biztalkserverreceiver/biztalkserverreceiver/internal/metadata"
	"go.opentelemetry.io/collector/scraper"
	"go.opentelemetry.io/collector/scraper/scraperhelper"
	"go.uber.org/zap"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

var (
	typeStr = component.MustNewType("biztalkserver")
)

const (
	defaultInterval = 1 * time.Minute
)

func createDefaultConfig() component.Config {
	return &Config{
		Endpoint: "http://localhost/biztalk/",
		Interval: string(defaultInterval),
	}
}

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		typeStr,
		createDefaultConfig,
		receiver.WithMetrics(createMetricsReceiver, metadata.MetricsStability))
}

func setupScraper(logger *zap.Logger, config *Config, settings receiver.Settings) *biztalkservermetricsScraper {
	return &biztalkservermetricsScraper{
		logger: logger,
		config: config,
		mb:     metadata.NewMetricsBuilder(config.MetricsBuilderConfig, settings),
	}
}

func createMetricsReceiver(
	_ context.Context,
	params receiver.Settings,
	conf component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	smrCfg := conf.(*Config)

	biztalkScraper := setupScraper(params.Logger, smrCfg, params)
	sc, err := scraper.NewMetrics(
		biztalkScraper.scrape,
		scraper.WithStart(biztalkScraper.Start),
		scraper.WithShutdown(biztalkScraper.Shutdown),
	)
	if err != nil {
		return nil, err
	}

	return scraperhelper.NewMetricsController(
		&smrCfg.ControllerConfig,
		params,
		consumer,
		scraperhelper.AddScraper(metadata.Type, sc),
	)
}

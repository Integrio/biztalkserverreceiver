package biztalkserverreceiver

import (
	"context"
	"github.com/Integrio/biztalkserverreceiver/biztalkserverreceiver/internal/metadata"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper"
	"go.opentelemetry.io/collector/scraper/scraperhelper"
	"time"
)

var (
	typeStr = component.MustNewType("biztalkserver")
)

const (
	defaultInterval = 1 * time.Minute
)

func createDefaultConfig() component.Config {
	cfg := scraperhelper.NewDefaultControllerConfig()
	cfg.CollectionInterval = defaultInterval

	return &Config{
		Endpoint:             "http://localhost/biztalk/",
		ControllerConfig:     cfg,
		MetricsBuilderConfig: metadata.DefaultMetricsBuilderConfig(),
	}
}

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		typeStr,
		createDefaultConfig,
		receiver.WithMetrics(createMetricsReceiver, metadata.MetricsStability))
}

func createMetricsReceiver(
	_ context.Context,
	params receiver.Settings,
	conf component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	smrCfg := conf.(*Config)

	biztalkScraper := newScraper(params.Logger, smrCfg, params)
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

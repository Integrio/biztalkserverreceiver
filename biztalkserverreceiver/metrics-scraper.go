package biztalkserverreceiver

import (
	"context"
	"fmt"
	"github.com/Integrio/biztalk-server-go/client"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
	"time"
)

type biztalkservermetricsScraper struct {
	host          component.Host
	cancel        context.CancelFunc
	logger        *zap.Logger
	nextConsumer  consumer.Metrics
	config        *Config
	simpleGauge   int64
	simpleCounter int64
	client        *client.Client
}

func (smr *biztalkservermetricsScraper) scrapeMetrics(ctx context.Context, metrics *pmetric.ResourceMetrics) {
	err := smr.scrapeOrchestrations(ctx, metrics)
	if err != nil {
		smr.logger.Error("Failed to scrape metrics", zap.Error(err))
	}
}

func (smr *biztalkservermetricsScraper) Start(ctx context.Context, host component.Host) error {
	smr.host = host
	ctx = context.Background()
	ctx, smr.cancel = context.WithCancel(ctx)
	interval, _ := time.ParseDuration(smr.config.Interval)
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				smr.logger.Info("Start processing metrcis")
				//ts := pcommon.NewTimestampFromTime(time.Now())

				metrics := pmetric.NewMetrics()
				resourceMetrics := metrics.ResourceMetrics().AppendEmpty()

				smr.scrapeMetrics(ctx, &resourceMetrics)

				err := smr.nextConsumer.ConsumeMetrics(ctx, metrics)
				if err != nil {
					smr.logger.Error(err.Error())
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}

func (smr *biztalkservermetricsScraper) Shutdown(ctx context.Context) error {
	if smr.cancel != nil {
		smr.cancel()
	}
	return nil
}

func (smr *biztalkservermetricsScraper) scrapeOrchestrations(ctx context.Context, rm *pmetric.ResourceMetrics) error {
	orchestrations, err := smr.client.GetOrchestrations(ctx)
	if err != nil {
		smr.logger.Error(err.Error())
	}

	scopeMetrics := rm.ScopeMetrics().AppendEmpty()

	for _, orch := range orchestrations {
		metric := scopeMetrics.Metrics().AppendEmpty()
		err := recordOrchestrationMetrics(&metric, &orch)
		if err != nil {
			smr.logger.Error(err.Error())
		}
	}

	return nil
}

func recordOrchestrationMetrics(metric *pmetric.Metric, orch *client.Orchestration) error {
	metric.SetName("biztalk.orchestrations_status")
	metric.SetDescription("Status of orchestrations in BizTalk Server.")
	metric.SetUnit("1")
	metric.SetEmptyGauge()

	var status int64
	switch orch.Status {
	case "Started":
		status = 1
		break
	case "Unenlisted":
		status = 0
		break
	case "Enlisted":
		status = 2
		break
	default:
		return fmt.Errorf("orchestration %s has invalid status", orch.Status)
	}

	dataPoint := metric.Gauge().DataPoints().AppendEmpty()
	dataPoint.SetTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	dataPoint.SetIntValue(status)

	dataPoint.Attributes().PutStr("full_name", orch.FullName)
	dataPoint.Attributes().PutStr("assembly_name", orch.AssemblyName)
	dataPoint.Attributes().PutStr("application_name", orch.AssemblyName)
	dataPoint.Attributes().PutStr("description", orch.Description)
	dataPoint.Attributes().PutStr("status", orch.Status)
	dataPoint.Attributes().PutStr("host", orch.Host)

	return nil
}

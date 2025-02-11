package biztalkserverreceiver

import (
	"context"
	errors2 "errors"
	"fmt"
	"github.com/Integrio/biztalk-server-go/client"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
	"strconv"
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

func (smr *biztalkservermetricsScraper) scrapeMetrics(ctx context.Context, metrics *pmetric.ResourceMetrics) error {
	errors := make([]error, 0)

	err := smr.scrapeOrchestrations(ctx, metrics)
	if err != nil {
		errors = append(errors, err)
		smr.logger.Error("Failed to scrape metrics", zap.Error(err))
	}

	err = smr.scrapeReceiveLocations(ctx, metrics)
	if err != nil {
		errors = append(errors, err)
		smr.logger.Error("Failed to scrape metrics", zap.Error(err))
	}

	err = smr.scrapeSendPorts(ctx, metrics)
	if err != nil {
		errors = append(errors, err)
		smr.logger.Error("Failed to scrape metrics", zap.Error(err))
	}

	err = smr.scrapeSendPortGroups(ctx, metrics)
	if err != nil {
		errors = append(errors, err)
		smr.logger.Error("Failed to scrape metrics", zap.Error(err))
	}

	err = smr.scrapeHostInstances(ctx, metrics)
	if err != nil {
		errors = append(errors, err)
		smr.logger.Error("Failed to scrape metrics", zap.Error(err))
	}

	return errors2.Join(errors...)
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
	case "Unenlisted":
		status = 0
		break
	case "Started":
		status = 1
		break
	case "Enlisted":
		status = 2
		break
	default:
		return fmt.Errorf("orchestration %s has invalid status %s", orch.FullName, orch.Status)
	}

	dp := metric.Gauge().DataPoints().AppendEmpty()
	dp.SetTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	dp.SetIntValue(status)

	dp.Attributes().PutStr("orchestration.full_name", orch.FullName)
	dp.Attributes().PutStr("orchestration.status", orch.Status)
	dp.Attributes().PutStr("orchestration.description", orch.Description)
	dp.Attributes().PutStr("orchestration.host", orch.Host)
	dp.Attributes().PutStr("assembly_name", orch.AssemblyName)
	dp.Attributes().PutStr("application_name", orch.AssemblyName)

	return nil
}

func (smr *biztalkservermetricsScraper) scrapeReceiveLocations(ctx context.Context, rm *pmetric.ResourceMetrics) error {
	receiveLocations, err := smr.client.GetReceiveLocations(ctx)
	if err != nil {
		smr.logger.Error(err.Error())
	}

	scopeMetrics := rm.ScopeMetrics().AppendEmpty()

	for _, rl := range receiveLocations {
		metric := scopeMetrics.Metrics().AppendEmpty()
		err := recordReceiveLocationMetrics(&metric, &rl)
		if err != nil {
			smr.logger.Error(err.Error())
		}
	}

	return nil
}

func recordReceiveLocationMetrics(metric *pmetric.Metric, rl *client.ReceiveLocation) error {
	metric.SetName("biztalk.receive_locations_status")
	metric.SetDescription("Enable status of receive locations in BizTalk Server.")
	metric.SetUnit("1")
	metric.SetEmptyGauge()

	dp := metric.Gauge().DataPoints().AppendEmpty()
	dp.SetTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	dp.SetIntValue(int64(boolToInt(rl.Enable)))

	dp.Attributes().PutStr("receive_location.name", rl.Name)
	dp.Attributes().PutStr("receive_location.status", strconv.FormatBool(rl.Enable))
	dp.Attributes().PutStr("receive_location.description", rl.Description)

	return nil
}

func (smr *biztalkservermetricsScraper) scrapeSendPorts(ctx context.Context, rm *pmetric.ResourceMetrics) error {
	sendPorts, err := smr.client.GetSendPorts(ctx)
	if err != nil {
		smr.logger.Error(err.Error())
	}

	scopeMetrics := rm.ScopeMetrics().AppendEmpty()

	for _, sp := range sendPorts {
		metric := scopeMetrics.Metrics().AppendEmpty()
		err := recordSendPortMetrics(&metric, &sp)
		if err != nil {
			smr.logger.Error(err.Error())
		}
	}

	return nil
}

func recordSendPortMetrics(metric *pmetric.Metric, sp *client.SendPort) error {
	metric.SetName("biztalk.send_ports_status")
	metric.SetDescription("Status of send ports in BizTalk Server.")
	metric.SetUnit("1")
	metric.SetEmptyGauge()

	var status int64
	switch sp.Status {
	case "Bound":
		status = 0
		break
	case "Started":
		status = 1
		break
	case "Stopped":
		status = 2
		break
	default:
		return fmt.Errorf("send port %s has invalid status %s", sp.Name, sp.Status)
	}

	dp := metric.Gauge().DataPoints().AppendEmpty()
	dp.SetTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	dp.SetIntValue(status)

	dp.Attributes().PutStr("send_port.name", sp.Name)
	dp.Attributes().PutStr("send_port.status", sp.Status)
	dp.Attributes().PutStr("send_port.description", sp.Description)
	dp.Attributes().PutStr("application_name", sp.ApplicationName)

	return nil
}

func (smr *biztalkservermetricsScraper) scrapeSendPortGroups(ctx context.Context, rm *pmetric.ResourceMetrics) error {
	spGroups, err := smr.client.GetSendPortGroups(ctx)
	if err != nil {
		smr.logger.Error(err.Error())
	}

	scopeMetrics := rm.ScopeMetrics().AppendEmpty()

	for _, spg := range spGroups {
		metric := scopeMetrics.Metrics().AppendEmpty()
		err := recordSendPortGroupMetrics(&metric, &spg)
		if err != nil {
			smr.logger.Error(err.Error())
		}
	}

	return nil
}

func recordSendPortGroupMetrics(metric *pmetric.Metric, spg *client.SendPortGroup) error {
	metric.SetName("biztalk.send_port_group_status")
	metric.SetDescription("Status of send port groups in BizTalk Server.")
	metric.SetUnit("1")
	metric.SetEmptyGauge()

	var status int64
	switch spg.Status {
	case "Bound":
		status = 0
		break
	case "Started":
		status = 1
		break
	case "Stopped":
		status = 2
		break
	default:
		return fmt.Errorf("send port group %s has invalid status %s", spg.Name, spg.Status)
	}

	dp := metric.Gauge().DataPoints().AppendEmpty()
	dp.SetTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	dp.SetIntValue(status)

	dp.Attributes().PutStr("send_port_group.name", spg.Name)
	dp.Attributes().PutStr("send_port_group.status", spg.Status)
	dp.Attributes().PutStr("send_port_group.description", spg.Description)
	dp.Attributes().PutStr("application_name", spg.ApplicationName)

	return nil
}

func (smr *biztalkservermetricsScraper) scrapeHostInstances(ctx context.Context, rm *pmetric.ResourceMetrics) error {
	hostInstances, err := smr.client.GetHostInstances(ctx)
	if err != nil {
		smr.logger.Error(err.Error())
	}

	scopeMetrics := rm.ScopeMetrics().AppendEmpty()

	for _, hi := range hostInstances {
		metric := scopeMetrics.Metrics().AppendEmpty()
		err := recordHostInstanceMetrics(&metric, &hi)
		if err != nil {
			smr.logger.Error(err.Error())
		}
	}

	return nil
}

func recordHostInstanceMetrics(metric *pmetric.Metric, hi *client.HostInstance) error {
	metric.SetName("biztalk.host_instance_status")
	metric.SetDescription("Status of host instances in BizTalk Server.")
	metric.SetUnit("1")
	metric.SetEmptyGauge()

	var status int64
	var statusStr string
	switch hi.IsDisabled {
	case true:
		statusStr = "Stopped"
		status = 0
		break
	case false:
		statusStr = "Running"
		status = 2
		break
	default:
		statusStr = "Unknown" // TODO: When should this be unknown? https://github.com/Integrio/biztalkserverreceiver/blob/64140539955b7fda174e824d8eacebd29d2eeb81/README.md?plain=1#L271
		status = 1
		break
	}

	dp := metric.Gauge().DataPoints().AppendEmpty()
	dp.SetTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	dp.SetIntValue(status)

	dp.Attributes().PutStr("host_instance.name", hi.Name)
	dp.Attributes().PutStr("host_instance.status", statusStr)
	dp.Attributes().PutStr("host_instance.service_state", hi.ServiceState)
	dp.Attributes().PutStr("host_instance.host_name", hi.HostName)
	dp.Attributes().PutStr("host_instance.running_server", hi.RunningServer)

	return nil
}

func boolToInt(b bool) int {
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return i
}

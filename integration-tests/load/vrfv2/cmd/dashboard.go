package main

import (
	"os"

	"github.com/K-Phoen/grabana/dashboard"
	"github.com/K-Phoen/grabana/logs"
	"github.com/K-Phoen/grabana/row"
	"github.com/K-Phoen/grabana/target/prometheus"
	"github.com/K-Phoen/grabana/timeseries"
	"github.com/K-Phoen/grabana/timeseries/axis"
	"github.com/smartcontractkit/wasp"
)

func main() {
	lokiDS := os.Getenv("DATA_SOURCE_NAME")
	d, err := wasp.NewDashboard(nil,
		[]dashboard.Option{
			dashboard.Row("LoadContractMetrics",
				row.WithTimeSeries(
					"RequestCount + FulfilmentCount",
					timeseries.Span(12),
					timeseries.Height("300px"),
					timeseries.DataSource(lokiDS),
					timeseries.Axis(
						axis.Unit("Requests"),
					),
					timeseries.WithPrometheusTarget(
						`
					last_over_time({type="vrfv2_contracts_load_summary", network=~"${network:pipe}", test_type=~"${test_type:pipe}", run_id=~"${run_id:pipe}", go_test_name=~"${go_test_name:pipe}", branch=~"${branch:pipe}", commit=~"${commit:pipe}", gen_name=~"${gen_name:pipe}"}
					| json
					| unwrap RequestCount [$__interval]) by (node_id, go_test_name, gen_name, run_id, test_type, network)
					`, prometheus.Legend("{{go_test_name}} requests"),
					),
					timeseries.WithPrometheusTarget(
						`
					last_over_time({type="vrfv2_contracts_load_summary", network=~"${network:pipe}", test_type=~"${test_type:pipe}", run_id=~"${run_id:pipe}", go_test_name=~"${go_test_name:pipe}", branch=~"${branch:pipe}", commit=~"${commit:pipe}", gen_name=~"${gen_name:pipe}"}
					| json
					| unwrap FulfilmentCount [$__interval]) by (node_id, go_test_name, gen_name, run_id, test_type, network)
					`, prometheus.Legend("{{go_test_name}} fulfillments"),
					),
				),
				row.WithTimeSeries(
					"Fulfillment time (blocks)",
					timeseries.Span(12),
					timeseries.Height("300px"),
					timeseries.DataSource(lokiDS),
					timeseries.Axis(
						axis.Unit("Blocks"),
					),
					timeseries.WithPrometheusTarget(
						`
					last_over_time({type="vrfv2_contracts_load_summary", network=~"${network:pipe}", test_type=~"${test_type:pipe}", run_id=~"${run_id:pipe}", go_test_name=~"${go_test_name:pipe}", branch=~"${branch:pipe}", commit=~"${commit:pipe}", gen_name=~"${gen_name:pipe}"}
					| json
					| unwrap AverageFulfillmentInMillions [$__interval]) by (node_id, go_test_name, gen_name, run_id, test_type, network) / 1e6
					`, prometheus.Legend("{{go_test_name}} avg"),
					),
					timeseries.WithPrometheusTarget(
						`
					last_over_time({type="vrfv2_contracts_load_summary", network=~"${network:pipe}", test_type=~"${test_type:pipe}", run_id=~"${run_id:pipe}", go_test_name=~"${go_test_name:pipe}", branch=~"${branch:pipe}", commit=~"${commit:pipe}", gen_name=~"${gen_name:pipe}"}
					| json
					| unwrap SlowestFulfillment [$__interval]) by (node_id, go_test_name, gen_name, run_id, test_type, network)
					`, prometheus.Legend("{{go_test_name}} slowest"),
					),
					timeseries.WithPrometheusTarget(
						`
					last_over_time({type="vrfv2_contracts_load_summary", network=~"${network:pipe}", test_type=~"${test_type:pipe}", run_id=~"${run_id:pipe}", go_test_name=~"${go_test_name:pipe}", branch=~"${branch:pipe}", commit=~"${commit:pipe}", gen_name=~"${gen_name:pipe}"}
					| json
					| unwrap FastestFulfillment [$__interval]) by (node_id, go_test_name, gen_name, run_id, test_type, network)
					`, prometheus.Legend("{{go_test_name}} fastest"),
					),
				),
			),
			dashboard.Row("CL nodes logs",
				row.Collapse(),
				row.WithLogs(
					"CL nodes logs",
					logs.DataSource(lokiDS),
					logs.Span(12),
					logs.Height("300px"),
					logs.Transparent(),
					logs.WithLokiTarget(`
					{type="log_watch"}
				`),
				)),
		},
	)
	if err != nil {
		panic(err)
	}
	// set env vars
	//export GRAFANA_URL=...
	//export GRAFANA_TOKEN=...
	//export DATA_SOURCE_NAME=Loki
	//export DASHBOARD_FOLDER=LoadTests
	//export DASHBOARD_NAME=WaspVRFv2
	if _, err := d.Deploy(); err != nil {
		panic(err)
	}
}

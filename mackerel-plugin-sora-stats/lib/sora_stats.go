package mpsorastats

import (
  "flag"
  "fmt"
  "io"

  "encoding/json"
  "net/http"

  mp "github.com/mackerelio/go-mackerel-plugin"
)

// JSON struct for sora stats API response
type SoraStats struct {
  AverageDurationSec         int `json:"average_duration_sec"`
  AverageSetupTimeMsec       int `json:"average_setup_time_msec"`
  TotalDurationSec           int `json:"total_duration_sec"`
  TotalFailedConnections     int `json:"total_failed_connections"`
  TotalOngoingConnections    int `json:"total_ongoing_connections"`
  TotalSuccessfulConnections int `json:"total_successful_connections"`
  TotalTurnTcpConnections    int `json:"total_turn_tcp_connections"`
  TotalTurnUdpConnections    int `json:"total_turn_udp_connections"`
}

// SorastatsPlugin mackerel plugin
type SorastatsPlugin struct {
  URI    string
  Prefix string
}

// MetricKeyPrefix interface for PluginWithPrefix
func (s SorastatsPlugin) MetricKeyPrefix() string {
  if s.Prefix == "" {
    s.Prefix = "sora-stats"
  }
  return s.Prefix
}

// GraphDefinition interface for mackerelplugin
func (s SorastatsPlugin) GraphDefinition() map[string]mp.Graphs {
  return map[string]mp.Graphs{
    "ongoing_connections": {
      Label: "SORA Ongoing Connections",
      Unit: mp.UnitInteger,
      Metrics: []mp.Metrics{
        {Name: "ongoing_connections", Label: "Current Ongoing connections", Diff: false},
      },
    },
    "average_connections": {
      Label: "SORA Average Connections in 1 minutes",
      Unit: mp.UnitFloat,
      Metrics: []mp.Metrics{
        {Name: "successful_connections", Label: "Successful connections in last 1 minutes", Diff: true},
        {Name: "failed_connections", Label: "Failed connections in last 1 minutes", Diff: true},
      },
    },
    "duration": {
      Label: "SORA Average Duration [sec]",
      Unit: mp.UnitInteger,
      Metrics: []mp.Metrics{
        {Name: "average_duration_sec", Label: "Average duration [sec]", Diff: false},
      },
    },
    "setup_time": {
      Label: "SORA Average Setup Time [msec]",
      Unit: mp.UnitInteger,
      Metrics: []mp.Metrics{
        {Name: "average_setup_time_msec", Label: "Average setup time [msec]", Diff: false},
      },
    },
    "turn_connections": {
      Label: "SORA Turn Connections in 1 minutes",
      Unit: mp.UnitFloat,
      Metrics: []mp.Metrics{
        {Name: "turn_tcp_connections", Label: "turn tcp connections in last 1 minutes", Diff: true},
        {Name: "turn_udp_connections", Label: "turn udp connections in last 1 minutes", Diff: true},
      },
    },
  }
}

// FetchMetrics interface for mackerelplugin
func (s SorastatsPlugin) FetchMetrics() (map[string]float64, error) {
  // Fetch stats report from SORA
  req, err := http.NewRequest("POST", s.URI, nil)
  if err != nil {
    return nil, err
  }
  req.Header.Set("x-sora-target", "Sora_20171010.GetStatsReport")
  res, err := http.DefaultClient.Do(req)
  if err != nil {
    return nil, err
  } else if res.StatusCode != 200 {
    return nil, fmt.Errorf("Unable to get stats by status code (%s) from url (%s)", res.StatusCode, s.URI)
  }
  defer res.Body.Close()

  return s.parseStats(res.Body)
}

// Parse sora stats API response body
func (s SorastatsPlugin) parseStats(body io.Reader) (map[string]float64, error) {
  stats := make(map[string]float64)

  // decode to json
  var soraStats SoraStats
  if err := json.NewDecoder(body).Decode(&soraStats); err == nil {
    stats["ongoing_connections"] = float64(soraStats.TotalOngoingConnections)
    stats["successful_connections"] = float64(soraStats.TotalSuccessfulConnections)
    stats["failed_connections"] = float64(soraStats.TotalFailedConnections)
    stats["average_duration_sec"] = float64(soraStats.AverageDurationSec)
    stats["average_setup_time_msec"] = float64(soraStats.AverageSetupTimeMsec)
    stats["turn_tcp_connections"] = float64(soraStats.TotalTurnTcpConnections)
    stats["turn_udp_connections"] = float64(soraStats.TotalTurnUdpConnections)
  } else {
    return nil, err
  }

  return stats, nil
}

// Do the plugin
func Do() {
  optURI := flag.String("uri", "", "URI")
  optScheme := flag.String("scheme", "http", "Scheme")
  optHost := flag.String("host", "localhost", "Hostname")
  optPort := flag.String("port", "port", "Port")
  optPrefix := flag.String("metric-key-prefix", "", "Metric key prefix")
  optTempfile := flag.String("tempfile", "", "Temp file name")
  flag.Parse()

  var sora SorastatsPlugin
  if *optURI != "" {
    sora.URI = *optURI
  } else {
    sora.URI = fmt.Sprintf("%s://%s:%s", *optScheme, *optHost, *optPort)
  }
  sora.Prefix = *optPrefix

  helper := mp.NewMackerelPlugin(sora)
  helper.Tempfile = *optTempfile
  helper.Run()
}

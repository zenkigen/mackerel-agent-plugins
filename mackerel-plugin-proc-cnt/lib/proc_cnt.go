package mpproccnt

import (
  "flag"
  "fmt"
  "os"
  "regexp"
  "strconv"
  "strings"

  mp "github.com/mackerelio/go-mackerel-plugin"
  "github.com/mattn/go-pipeline"
)

// ProccntPlugin mackerel plugin
type ProccntPlugin struct {
  Process           string
  Prefix            string
  NormalizedProcess string
  MetricName        string
}

// MetricKeyPrefix interface for PluginWithPrefix
func (p ProccntPlugin) MetricKeyPrefix() string {
  if p.Prefix == "" {
    p.Prefix = "proc-cnt"
  }
  return p.Prefix
}

// GraphDefinition interface for mackerelplugin
func (p ProccntPlugin) GraphDefinition() map[string]mp.Graphs {
  return map[string]mp.Graphs{
    p.NormalizedProcess: {
      Label: fmt.Sprintf("Process Count of %s", p.NormalizedProcess),
      Unit: mp.UnitInteger,
      Metrics: []mp.Metrics{
        {Name: "processes", Label: "Active processes", Diff: false},
      },
    },
  }
}

// FetchMetrics interface for mackerelplugin
func (p ProccntPlugin) FetchMetrics() (map[string]float64, error) {
  stats := make(map[string]float64)
  // Fetch all pids withc contains specified process name
  out, err := pipeline.Output(
     []string{"ps", "aux"},
     []string{"grep", p.Process},
     []string{"grep", "-v", "grep"},
     []string{"grep", "-v", "mackerel-plugin-proc-cnt"},
     []string{"wc", "-l"},
  )

  if err != nil {
    // No matching with p.Process
    stats["processes"] = 0
    return stats, nil
  }

  proc_num, err := strconv.ParseFloat(strings.TrimSpace(string(out)), 64)

  if err != nil {
    return nil, err
  }

  stats["processes"] = proc_num
  return stats, nil
}

func normalizeForMetricName(process string) string {
  // Mackerel accepts following characters in custom metric names
  // [-a-zA-Z0-9_.]
  re := regexp.MustCompile("[^-a-zA-Z0-9_.]")
  return re.ReplaceAllString(process, "_")
}

// Do the plugin
func Do() {
  optProcess := flag.String("process", "", "Process name")
  optPrefix := flag.String("metric-key-prefix", "", "Metric key prefix")
  optTempfile := flag.String("tempfile", "", "Temp file name")
  flag.Parse()

  if *optProcess == "" {
    flag.PrintDefaults()
    os.Exit(1)
  }

  var cnt ProccntPlugin
  cnt.Process = *optProcess
  cnt.Prefix = *optPrefix
  cnt.NormalizedProcess = normalizeForMetricName(*optProcess)

  helper := mp.NewMackerelPlugin(cnt)
  helper.Tempfile = *optTempfile
  helper.Run()
}

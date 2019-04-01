mackerel-plugin-proc-cnt
=======================
Process count custom metrics plugin for mackerel.io.agent.

## Synopsis
```shell
mackerel-plugin-proc-cnt [-process=<Process name>]
```

## Exapmle of mackerel-agent-conf
```
[plugin.metrics.proc_cnt]
command = "/path/to/mackerel-plugin-proc-cnt -process nginx"
```

## Build

```
$ go get github.com/zenkigen/mackerel-agent-plugins
$ github.com/mackerelio/go-mackerel-plugin
$ go get github.com/mattn/go-pipeline
$ go build main.go
```

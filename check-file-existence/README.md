mackerel-plugin-check-existence
=======================
File or Directory check plugin for mackerel.io.agent.

## Synopsis
```shell
check-file-existence [-path=<Path to file or directory>]
```

## Exapmle of mackerel-agent-conf
```
[plugin.metrics.proc_cnt]
command = "/path/to/check-file-existence -path /tmp/archive"
```

## Build

```
$ go get github.com/zenkigen/mackerel-agent-plugins
$ go get github.com/mackerelio/checkers
$ go get github.com/mattn/go-pipeline
$ go build main.go
```

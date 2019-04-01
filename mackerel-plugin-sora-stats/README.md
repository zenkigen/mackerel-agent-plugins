mackerel-plugin-sora-stats
=======================
Sora stats report API custom metrics plugin for mackerel.io.agent.

## Synopsis
```shell
mackerel-plugin-sora-sora [-host=<Hostname>] [-port=<Port>] [-scheme=<http|https>] [-tempfile=<Path to tmp file>] [-uri=<URI>]
```

## Exapmle of mackerel-agent-conf
```
[plugin.metrics.sora_stats]
command = "/path/to/mackerel-plugin-sora-sora"
```

## Build

```
$ go get github.com/zenkigen/mackerel-agent-plugins
$ go get github.com/mackerelio/go-mackerel-plugin
$ go build main.go
```

## Reference

### About Sora by Shiguredo
https://sora.shiguredo.jp

https://sora.shiguredo.jp/doc/index.html

### Sora Stats Report API
https://sora.shiguredo.jp/doc/API.html#id28

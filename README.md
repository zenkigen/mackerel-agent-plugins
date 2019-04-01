# mackerel-agent-plugins
This is the additional unofficial pack of mackerel-agent plugins.
The official mackerel-agent-plugins is refered to below.
https://github.com/mackerelio/mackerel-agent-plugins

## Plugins in this pack

* mackerel-plugin-check-existence
  - File or Directory check plugin by 'ls' command for mackerel.io.agent.
* mackerel-plugin-proc-cnt
  - Process count custom metrics plugin for mackerel.io.agent.
* mackerel-plugin-sora-stats
  - SORA stats report API custom metrics plugin for mackerel.io.agent.
  - SORA is a software package of WebRTC SFU released by Shiguredo.
  - refer to: https://sora.shiguredo.jp/

For detail, please look at each plugin's README.

# Install

Currently, the package in this repository is only for linux amd64.
If you want to use in other architecture, please contact us or submit issues.

## Linux

```
$ wget https://github.com/zenkigen/mackerel-agent-plugins/raw/master/zenkigen-mackerel-agent-plugins_1.0.0_amd64.deb
$ dpkg -i zenkigen-mackerel-agent-plugins_1.0.0_amd64.deb
```

# Build

Please see each plugin's README.

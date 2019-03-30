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


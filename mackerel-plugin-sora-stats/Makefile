.PHONY: all
all:
	go build -o mackerel-plugin-sora-stats main.go

.PHONY: package
package:
	${GOPATH}/bin/gox --osarch "linux/amd64"

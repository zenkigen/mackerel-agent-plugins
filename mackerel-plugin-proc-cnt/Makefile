.PHONY: all
all:
	go build -o mackerel-plugin-proc-cnt main.go

.PHONY: package
package:
	${GOPATH}/bin/gox --osarch "linux/amd64"

.PHONY: all
all:
	go build -o check-file-existence main.go

.PHONY: package
package:
	${GOPATH}/bin/gox --osarch "linux/amd64"

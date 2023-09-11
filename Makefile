help:
	@echo "help"
	@echo
	@echo "manuf: Pulls latest /etc/manuf from www.wireshark.org"
	@echo "build: Compiles and links in current git commit and timestamp"
	@echo "install: go install with current git commit and timestamp"
	@echo "clean: Delete build artifacts"

test:	manuf
	go test -v ./...

manuf:
	wget -O manuf "https://www.wireshark.org/download/automated/data/manuf"

build:
	go build ./...

install:
	go install ./...

clean:
	rm -f manuf
	rm -f plausible
	rm -f $$GOPATH/bin/plausible

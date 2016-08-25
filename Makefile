BUILDFLAGS = -ldflags="-X main.commit=`git rev-parse HEAD` -X main.buildTime=`date -u '+%Y-%m-%d_%I:%M:%S%p'`"

help:
	@echo "help"
	@echo
	@echo "manuf: Pulls latest /etc/manuf from code.wireshark.org"
	@echo "build: Compiles and links in current git commit and timestamp"
	@echo "install: go install"
manuf:
	wget -O manuf "https://code.wireshark.org/review/gitweb?p=wireshark.git;a=blob_plain;f=manuf"

build:
	go build $(BUILDFLAGS)

install:
	go install $(BUILDFLAGS)

clean:
	rm manuf
	rm plausible

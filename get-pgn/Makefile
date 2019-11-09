# Variable section
BINARY := get-pgn
.DEFAULT_GOAL := release
PLATFORMS := windows linux darwin
GOBUILD=go build
BINDIR=../bin/
GOTGT=-o $(BINDIR)$@ .
os = $(word 1, $@)

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	mkdir -p ../bin
	GOOS=$(os) GOARCH=amd64 go build -o ../bin/$(BINARY)-$(os)-amd64

.PHONY: release
release: windows linux darwin

.PHONY: clean
clean:
	-cd ../bin;rm -rf get-pgn get-pgn.darwin get-pgn.linux get-pgn.exe

.PHONY: clean-all
clean-all: clean
	-rm -rf *.pgn


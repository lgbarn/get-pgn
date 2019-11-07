# Variable section
GOBUILD=go build
BINDIR=../bin/
GOTGT=-o $(BINDIR)$@ .

.PHONY: all
all: $(BINDIR)get-pgn.darwin $(BINDIR)get-pgn.linux $(BINDIR)get-pgn.exe

$(BINDIR)get-pgn.darwin: main.go
	GOOS=darwin ${GOBUILD} ${GOTGT}

$(BINDIR)get-pgn.linux: main.go
	GOOS=linux ${GOBUILD} ${GOTGT}

$(BINDIR)get-pgn.exe: main.go
	GOOS=windows ${GOBUILD} ${GOTGT}

.PHONY: clean
clean:
	-cd ../bin;rm -rf get-pgn get-pgn.darwin get-pgn.linux get-pgn.exe

.PHONY: clean-all
clean-all: clean
	-rm -rf *.pgn


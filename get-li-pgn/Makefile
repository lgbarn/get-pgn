# Variable section
GOBUILD=go build
BINDIR=../bin/
GOTGT=-o $(BINDIR)$@ .

.PHONY: all
all: $(BINDIR)get-li-pgn.darwin $(BINDIR)get-li-pgn.linux $(BINDIR)get-li-pgn.exe

$(BINDIR)get-li-pgn.darwin: main.go
	GOOS=darwin ${GOBUILD} ${GOTGT}

$(BINDIR)get-li-pgn.linux: main.go
	GOOS=linux ${GOBUILD} ${GOTGT}

$(BINDIR)get-li-pgn.exe: main.go
	GOOS=windows ${GOBUILD} ${GOTGT}

.PHONY: clean
clean:
	-cd ../bin; rm -rf get-li-pgn get-li-pgn.darwin get-li-pgn.linux get-li-pgn.exe

.PHONY: clean-all
clean-all: clean
	-rm -rf *.pgn


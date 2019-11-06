# Variable section
GOBUILD=go build
GOTGT=-o $@ .

.PHONY: default
default:
	${GOBUILD} .

.PHONY: all
all: get-pgn.darwin get-pgn.linux get-pgn.exe

get-pgn.darwin: main.go
	GOOS=darwin ${GOBUILD} ${GOTGT}

get-pgn.linux: main.go
	GOOS=linux ${GOBUILD} ${GOTGT}

get-pgn.exe: main.go
	GOOS=windows ${GOBUILD} ${GOTGT}

.PHONY: clean
clean:
	-rm -rf get-pgn get-pgn.darwin get-pgn.linux get-pgn.exe

.PHONY: clean-all
clean-all: clean
	-rm -rf *.pgn


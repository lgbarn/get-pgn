# Variable section
GOBUILD=go build
GOTGT=-o $@ .

.PHONY: default
default:
	${GOBUILD} .

.PHONY: all
all: get-li-pgn.darwin get-li-pgn.linux get-li-pgn.exe

get-li-pgn.darwin: main.go
	GOOS=darwin ${GOBUILD} ${GOTGT}

get-li-pgn.linux: main.go
	GOOS=linux ${GOBUILD} ${GOTGT}

get-li-pgn.exe: main.go
	GOOS=windows ${GOBUILD} ${GOTGT}

.PHONY: clean
clean:
	-rm -rf get-li-pgn get-li-pgn.darwin get-li-pgn.linux get-li-pgn.exe

.PHONY: clean-all
clean-all: clean
	-rm -rf *.pgn


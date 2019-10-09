# Variable section
GO=go build
GOTGT=-o $@ .


.PHONY: all
all: get-li-pgn.darwin get-li-pgn.linux get-li-pgn.exe

get-li-pgn.darwin: main.go
	GOOS=darwin ${GO} ${GOTGT}

get-li-pgn.linux: main.go
	GOOS=linux ${GO} ${GOTGT}

get-li-pgn.exe: main.go
	GOOS=windows ${GO} ${GOTGT}

.PHONY: clean
clean:
	-rm -rf get-li-pgn get-li-pgn.darwin get-li-pgn.linux get-li-pgn.exe

.PHONY: clean-all
clean-all: clean
	-rm -rf *.pgn


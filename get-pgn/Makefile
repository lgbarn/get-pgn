# Variable section
GO=go build
GOTGT=-o $@ .


.PHONY: all
all: get-pgn.darwin get-pgn.linux get-pgn.exe

get-pgn.darwin: main.go
	GOOS=darwin ${GO} ${GOTGT}

get-pgn.linux: main.go
	GOOS=linux ${GO} ${GOTGT}

get-pgn.exe: main.go
	GOOS=windows ${GO} ${GOTGT}

.PHONY: clean
clean:
	-rm -rf get-pgn get-pgn.darwin get-pgn.linux get-pgn.exe

.PHONY: clean-all
clean-all: clean
	-rm -rf *.pgn


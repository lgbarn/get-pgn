.PHONY : all
all:
	cd get-li-pgn && make $@
	cd get-pgn && make $@

.PHONY : clean
clean:
	cd get-li-pgn && make $@
	cd get-pgn && make $@


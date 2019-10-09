SUBDIRS = get-li-pgn get-pgn

.PHONY : all
all: ${SUBDIRS}

.PHONY : ${SUBDIRS}
${SUBDIRS}:
	${MAKE} -C $@

.PHONY : clean
clean:
	for dir in ${SUBDIRS}; do \
		${MAKE} -C $${dir} clean; \
	done

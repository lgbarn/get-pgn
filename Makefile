SUBDIRS = get-li-pgn get-pgn

.PHONY : default
default: 
	for dir in ${SUBDIRS}; do \
		${MAKE} -C $${dir} ; \
	done

.PHONY : all
all: 
	for dir in ${SUBDIRS}; do \
		${MAKE} -C $${dir} all; \
	done

.PHONY : ${SUBDIRS}
${SUBDIRS}:

.PHONY : clean
clean:
	for dir in ${SUBDIRS}; do \
		${MAKE} -C $${dir} clean; \
	done

.PHONY: lint
lint:
	golint ./...


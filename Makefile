OPTS=-v -cover
FIND_OPTS=-maxdepth 1 -type d -regextype posix-extended -regex '^\.\/problem[0-9]{1,3}$$'
TARGET=

FORCE:

test: FORCE
	[ -z '$(TARGET)' ] && find . $(FIND_OPTS) -exec go test $(OPTS) {} \;; [ -z '$(TARGET)' ] || go test $(OPTS) './$(TARGET)'

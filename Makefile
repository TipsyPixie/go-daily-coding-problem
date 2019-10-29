OPTS=-v -cover
FIND_OPTS=-maxdepth 1 -type d -regextype posix-extended -regex '^\.\/problem[0-9]{1,3}$$'
NO=

FORCE:

test: FORCE
	[ -z '$(NO)' ] && find . $(FIND_OPTS) -exec go test $(OPTS) {} \;; [ -z '$(NO)' ] || go test $(OPTS) './problem$(NO)'

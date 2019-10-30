TEST_OPTS=-v -cover
FIND_OPTS=-maxdepth 1 -type d -regextype posix-extended -regex '^\.\/problem[0-9]{1,3}$$'
VET_OPTS=
NO=

FORCE:

test: FORCE
	[ -z '$(NO)' ] && find . $(FIND_OPTS) -exec go test $(TEST_OPTS) {} \;; [ -z '$(NO)' ] || go test $(TEST_OPTS) './problem$(NO)'

analyze: FORCE
	[ -z '$(NO)' ] && find . $(FIND_OPTS) -exec go vet $(VET_OPTS) {} \;; [ -z '$(NO)' ] || go vet $(VET_OPTS) './problem$(NO)'

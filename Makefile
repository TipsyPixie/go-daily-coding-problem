TEST_OPTS=-v -cover
FIND=find -maxdepth 1 -type d -regextype posix-extended -regex '^\.\/problem[0-9]{1,3}$$'
VET_OPTS=
GOFMT_OPTS=-l -s -w
FIX_OPTS=
NO=

FORCE:

test: FORCE
	[ -z '$(NO)' ] && $(FIND) | sort | xargs go test $(TEST_OPTS); [ -z '$(NO)' ] || go test $(TEST_OPTS) './problem$(NO)'

analyze: FORCE
	[ -z '$(NO)' ] && $(FIND) | sort | xargs go vet $(VET_OPTS); [ -z '$(NO)' ] || go vet $(VET_OPTS) './problem$(NO)'

format: FORCE
	[ -z '$(NO)' ] && $(FIND) | sort | xargs gofmt $(GOFMT_OPTS); [ -z '$(NO)' ] || gofmt $(GOFMT_OPTS) './problem$(NO)'

fix: FORCE
	[ -z '$(NO)' ] && $(FIND) | sort | xargs go fix $(FIX_OPTS); [ -z '$(NO)' ] || go fix $(FIX_OPTS) './problem$(NO)'

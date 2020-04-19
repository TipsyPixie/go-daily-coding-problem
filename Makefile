TEST_OPTS=-v -cover -timeout 10s -count 1
FIND=find -maxdepth 1 -type d -regextype posix-extended -regex '^\.\/problem[0-9]{1,3}$$'
VET_OPTS=
FMT_OPTS=-l -s -w
FIX_OPTS=

FORCE:

test: FORCE
	! $(FIND) -exec go test $(TEST_OPTS) {} \; | grep -q 'FAIL'

analyze: FORCE
	$(FIND) -exec go vet $(VET_OPTS) {} \;

format: FORCE
	$(FIND) -exec gofmt $(FMT_OPTS) {} \;

fix: FORCE
	$(FIND) -exec go fix $(FIX_OPTS) {} \;

report: FORCE
	curl -s -d 'repo=github.com%2FTipsyPixie%2Fgo-daily-coding-problem' https://goreportcard.com/checks >/dev/null

precommit: fix analyze format test

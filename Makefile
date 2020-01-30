TEST_OPTS=-v -cover -timeout 15s -count 1
FIND=find -maxdepth 1 -type d -regextype posix-extended -regex '^\.\/problem[0-9]{1,3}$$'
VET_OPTS=
GOFMT_OPTS=-l -s -w
FIX_OPTS=

FORCE:

test: FORCE
	! $(FIND) | sort | xargs go test $(TEST_OPTS) | grep -q 'FAIL'

analyze: FORCE
	$(FIND) | sort | xargs go vet $(VET_OPTS)

format: FORCE
	$(FIND) | sort | xargs gofmt $(GOFMT_OPTS)

fix: FORCE
	$(FIND) | sort | xargs go fix $(FIX_OPTS)

report: FORCE
	curl -s -d 'repo=github.com%2FTipsyPixie%2Fgo-daily-coding-problem' https://goreportcard.com/checks >/dev/null

hook: FORCE
	find 'hooks' -maxdepth 1 -type f -exec ln -s '../../{}' -t './.git/hooks' \;

precommit: fix analyze format test

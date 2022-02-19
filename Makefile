TESTARGS?=-cover -timeout 20s
FIND=find -maxdepth 1 -type d -regextype posix-extended -regex '^\.\/problem[0-9]{1,3}$$'
PACKAGES=$$(go list ./...)

FORCE:

test: FORCE
	go test ${TESTARGS} ${PACKAGES}

check: FORCE
	go vet ${VETARGS} ${PACKAGES}

format: FORCE
	go fmt ${FMTARGS} ${PACKAGES}

fix: FORCE
	go fix ${FIXARGS} ${PACKAGES}

report: FORCE
	curl -s --max-time 10 -d 'repo=github.com%2FTipsyPixie%2Fgo-daily-coding-problem' https://goreportcard.com/checks >/dev/null

precommit: fix format check test

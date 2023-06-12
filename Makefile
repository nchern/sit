.PHONY: build
build: vet
	@go build ./...

.PHONY: install
install: test
	@go install ./...

.PHONY: vet
vet:
	@go vet ./...

.PHONY: test
test: build
	@go test -race ./...

.PHONY: open-issues-report
open-issues-report:
	@./bin/gen-open-issues-report.sh > issues.md

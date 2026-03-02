GO ?= CGO_ENABLED=0 go

.cpuname:
	@$(GO) test -bench=BenchmarkDetectCPU -benchtime=1ns ./cmd/cpuname | $(GO) run ./cmd/cpuname > .cpuname

BENCH_FILE := benches/$(shell $(GO) env GOOS)-$(shell $(GO) env GOARCH)-$(shell cat .cpuname 2>/dev/null || echo unknown).txt

.PHONY: help
help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "%-15s %s\n", $$1, $$2}'

.PHONY: all
all: tidy test ## Run tidy and test

.PHONY: clean
clean: ## Remove generated files
	$(GO) clean
	-rm -f \
		.cpuname \
		coverage.txt \
		coverage.xml \
		gl-code-quality-report.json \
		golangci-lint.json \
		govulncheck.sarif \
		junit.xml \
		test.log

.PHONY: bench
bench: ## Run benchmarks
	$(GO) test -run=^$$ -bench=Day..Part.$$ -benchmem

.PHONY: tidy
tidy: ## Format check and lint
	test -z "$$(gofmt -l .)"
	$(GO) vet
	$(GO) run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run
	@with_tests=$$(mktemp); \
	without_tests=$$(mktemp); \
	with_funcs=$$(mktemp); \
	without_funcs=$$(mktemp); \
	test_only_funcs=$$(mktemp); \
	trap 'rm -f "$$with_tests" "$$without_tests" "$$with_funcs" "$$without_funcs" "$$test_only_funcs"' EXIT; \
	$(GO) run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --enable-only unused --tests=true --out-format tab ./... > "$$with_tests" 2>/dev/null || true; \
	$(GO) run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --enable-only unused --tests=false --out-format tab ./... > "$$without_tests" 2>/dev/null || true; \
	LC_ALL=C grep -E 'unused[[:space:]]+func `.*` is unused$$' "$$with_tests" | sed -E 's/[[:space:]]+/ /g' | sort -u > "$$with_funcs" || true; \
	LC_ALL=C grep -E 'unused[[:space:]]+func `.*` is unused$$' "$$without_tests" | sed -E 's/[[:space:]]+/ /g' | sort -u > "$$without_funcs" || true; \
	comm -23 "$$without_funcs" "$$with_funcs" > "$$test_only_funcs"; \
	if [ -s "$$test_only_funcs" ]; then \
		echo "WARNING: functions only kept alive by tests:"; \
		cat "$$test_only_funcs"; \
	else \
		echo "No functions are only kept alive by tests."; \
	fi

.PHONY: test
test: ## Run all days part 1 and 2 but no additional tests
	$(GO) test -run=Day..Part.$$

$(BENCH_FILE): .cpuname $(wildcard *.go)
	@mkdir -p benches
	@echo "Running benchmarks and saving to $@..."
ifeq ($(shell $(GO) env GOOS),linux)
	@if [ -d /sys/devices/system/cpu/cpu0/cpufreq ]; then \
		SAVED_GOV=$$(cat /sys/devices/system/cpu/cpu0/cpufreq/scaling_governor); \
		echo "Setting CPU governor to performance mode..."; \
		for cpu in /sys/devices/system/cpu/cpu*/cpufreq/scaling_governor; do \
			echo performance | sudo tee $$cpu > /dev/null; \
		done; \
		$(GO) test -run=^$$ -bench=Day..Part.$$ -benchmem | tee $@; \
		echo "Restoring CPU governor to $$SAVED_GOV..."; \
		for cpu in /sys/devices/system/cpu/cpu*/cpufreq/scaling_governor; do \
			echo $$SAVED_GOV | sudo tee $$cpu > /dev/null; \
		done; \
	else \
		$(GO) test -run=^$$ -bench=Day..Part.$$ -benchmem | tee $@; \
	fi
else
	$(GO) test -run=^$$ -bench=Day..Part.$$ -benchmem | tee $@
endif

.PHONY: total
total: .cpuname ## Run benchmarks and show total runtime
	@BENCH_FILE="benches/$$($(GO) env GOOS)-$$($(GO) env GOARCH)-$$(cat .cpuname).txt"; \
	$(MAKE) --no-print-directory "$$BENCH_FILE" && awk -f total.awk < "$$BENCH_FILE"

.PHONY: total-nogc
total-nogc: ## Run benchmarks with GOGC=off and show total runtime
	@mkdir -p benches
	GOGC=off $(GO) test -run=^$$ -bench=Day..Part.$$ -benchmem | tee $(BENCH_FILE)
	@awk -f total.awk < $(BENCH_FILE)

.PHONY: sast
sast: coverage.xml gl-code-quality-report.json govulncheck.sarif junit.xml ## Generate GitLab CI reports

coverage.txt test.log &:
	-$(GO) test -coverprofile=coverage.txt -covermode count -short -v | tee test.log

junit.xml: test.log
	$(GO) run github.com/jstemmer/go-junit-report/v2@latest < $< > $@

coverage.xml: coverage.txt
	$(GO) run github.com/boumenot/gocover-cobertura@latest < $< > $@

gl-code-quality-report.json: golangci-lint.json
	$(GO) run github.com/banyansecurity/golint-convert@latest < $< > $@

golangci-lint.json:
	-$(GO) run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --out-format json > $@

govulncheck.sarif:
	$(GO) run golang.org/x/vuln/cmd/govulncheck@latest -format=sarif ./... > $@

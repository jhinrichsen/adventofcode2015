GO ?= CGO_ENABLED=0 go
CPU_NAME := $(shell $(GO) run ./cmd/cpuname)
BENCH_FILE := benches/$(shell $(GO) env GOOS)-$(shell $(GO) env GOARCH)-$(CPU_NAME).txt

.PHONY: all
all: lint test

.PHONY: clean
clean:
	rm -f README.pdf README.html

.PHONY: bench
bench:
	$(GO) test -bench=. -run="" -benchmem

.PHONY: lint
lint:
	$(GO) vet
	$(GO) run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run

.PHONY: test
test:
	$(GO) test -coverprofile=coverage.txt -covermode count gitlab.com/jhinrichsen/adventofcode2015
	$(GO) run github.com/boumenot/gocover-cobertura@latest < coverage.txt > coverage.xml

prof:
	$(GO) -bench=. -benchmem -memprofile mprofile.out -cpuprofile cprofile.out
	$(GO) pprof cpu.profile

# some asciidoc targets
.PHONY: doc
doc: README.html README.pdf

README.html: README.adoc
	asciidoctor $<

README.pdf: README.adoc
	asciidoctor-pdf -a allow-uri-read $<

$(BENCH_FILE): $(wildcard *.go)
	@echo "Running benchmarks and saving to $@..."
	@mkdir -p benches
	GOGC=off $(GO) test -run=^$$ -bench='^BenchmarkDay[0-9]+Part[12]$$' -benchmem | tee $@

.PHONY: total
total: $(BENCH_FILE)
	@awk -f total.awk < $(BENCH_FILE)

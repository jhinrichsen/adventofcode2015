bench:
	go test -run NONE -bench . -benchmem -timeout=30m

benchstat: benchmarks/go1.16.6_darwin_amd64.txt benchmarks/go1.17_darwin_amd64.txt 
	benchstat $<

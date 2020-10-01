prof:
	go test -bench=Day6 -benchmem -memprofile mprofile.out -cpuprofile cprofile.out
	go tool pprof cpu.profile
	# go tool pprof mem.profile

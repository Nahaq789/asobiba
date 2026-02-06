@run_concurrent_go:
  go run ./concurrent_go/main.go

@run_parallel_go:
  go run ./parallel_go/main.go

@run_parallel_rust:
  cd parallel_rust && cargo run

@run_035:
  cd 0.35 && go run ./...

@run_uuid:
  cd uuid && go run ./...

accrual:
	./cmd/accrual/accrual_darwin_arm64 -a "localhost:8181" -d "postgres://root:pass@postgres:5432/gophermart?sslmode=disable"

test:
	go test -v ./... -count=1

build: test
	go build -o ./cmd/gophermart/bin/gophermart ./cmd/gophermart

run: build
	accrual
	./bin/gophermart
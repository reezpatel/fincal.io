schema := "public"
grpc_dir := "fincal"
binary_name := `awk '/^module/{print $2}' go.mod`
main_package_path := "./"


set dotenv-required
set dotenv-load := true

sloc:
    @echo "`wc -l **/*.go` lines of code"

jet:
    jet -dsn=$PG_LOCAL_DB_URL -schema={{schema}} -path=./.gen

audit:
    go vet ./...
    go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000,-ST1001 ./...
    go run golang.org/x/vuln/cmd/govulncheck@latest ./...

protoc:
    mkdir -p {{grpc_dir}}
    mkdir -p ./.gen/openapiv2
    protoc  -I . --grpc-gateway_out ./{{grpc_dir}} --grpc-gateway_opt grpc_api_configuration=./service.yaml \
            --go_out=./{{grpc_dir}} \
            --go-grpc_out=./{{grpc_dir}} \
            --openapiv2_out ./.gen/openapiv2 --openapiv2_opt grpc_api_configuration=./service.yaml \
            proto/*

test:
    go test -v -race -buildvcs ./...

textWithCoverage:
    go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
    go tool cover -html=/tmp/coverage.out

build:
    go build -a -installsuffix cgo -o=/tmp/bin/{{binary_name}} {{main_package_path}}

run: build
    /tmp/bin/{{binary_name}}

up: build
    /tmp/bin/{{binary_name}} --migrate up

down: build
    /tmp/bin/{{binary_name}} --migrate down

watch:
    go run github.com/cosmtrek/air@v1.43.0 \
        --build.cmd "make build" --build.bin "/tmp/bin/{{binary_name}}" --build.delay "100" \
        --build.exclude_dir "" \
        --misc.clean_on_exit "true"

prod:
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags='-s' -o=/tmp/bin/linux_amd64/{{binary_name}} {{main_package_path}}
    upx -5 /tmp/bin/linux_amd64/{{binary_name}}
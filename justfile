default: build unit-test

build:
  go build -o google-monitor ./cmd/...

unit-test:
  go test -v -cover ./...

integration-tests-real-api:
  docker-compose up -d mongo
  go test -v -cover --tags="real_api integration_tests" ./...
  docker-compose down

integration-tests:
  docker-compose up -d mongo
  go test -v -cover --tags="integration_tests" ./...
  docker-compose down

lint:
  echo 'TODO'
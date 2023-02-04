# google-monitor

This repository contains basic monitoring service,
which simply makes request to google, log response in mongo and std out.

It uses `just` to orchestrate  simple run commands.

Simple run `just` to execute default build, unit-testing.

`just build` - to build binary.

`just unit-test` - to run unit tests with mock api.

`just integration-tests-real-api` - to run integration tests with mongo and google api.

`just integration-tests` - to run integration tests.

`docker-compose up` - to run mongo, mongo-express and go service(it builds go service from docker file).

`docker build -t {tag}` . - to build docker file. 

### TODO

 CI/CD, k8s, linting
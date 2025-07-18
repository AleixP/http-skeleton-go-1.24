# HTTP Skeleton Application

Use this skeleton application to quickly setup and start working on a new HTTP Service.
This application provides a basic API REST developed following a DDD approach.

## Requirements

1. Git
1. Make
2. Docker & Docker Compose

## Installation


1. Install the application running `make install` in the application directory.
2. You can go to `http://localhost:8080/healthz` to check the service status.
3. If you have another service running at port 8080, you can run `make dcp` in order to clean all docker containers, and then rerun `make install`
4. Finally you can execute `make test` to run all test suites.

### Make

We are using **Make** to automate the most common development tasks. You can type `make` to see all available targets:

```
  install                    Install required software and initialize your local configuration
  build                      Build application containers
  test                       Run all test suites
  run                        Executes the app locally
```

## Testing

### Integration and Acceptance Tests

In order to help us produce and run integration and acceptance tests I am using [Testify](https://github.com/stretchr/testify/suite).
To run all your test suites just run:

```
$ make test
```

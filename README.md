# Golang Git lab/hub CI/CD Demo

A simple project to demonstrate how to use Gitlab CI/CD with Go.


## Developing

Install dependencies with:

```bash
go get -v -d ./...
```

Run the scripts with:

```bash
go run main.go
```

Run the tests with:

```bash
go test -v ./...
```


## Build

Use the CI file `.gitlab-ci.yml` to build the binary:


```yml
image: golang:latest

variables:
  REPO: my.gitlab.com
  GROUP: server_management
  PROJECT: go-hello

stages:
 - test
 - build

before_script:
  - mkdir -p $GOPATH/src/$REPO/$GROUP $GOPATH/src/_/builds
  - cp -r $CI_PROJECT_DIR $GOPATH/src/$REPO/$GROUP/$PROJECT
  - ln -s $GOPATH/src/$REPO/$GROUP $GOPATH/src/_/builds/$GROUP
  - go get -v -d ./...

unit_tests:
  stage: test
  script:
    - go test -v ./...

build:
  stage: build
  script:
    - go build -v -o hi
  only:
    - main
  artifacts:
    paths:
      - /builds/server_management/hello-go/hi
    expire_in: 1 hour
```

Download the build binary from:
    

> `https://my.gitlab.com/server_management/hello-go/-/jobs/artifacts/main/browse?job=build`


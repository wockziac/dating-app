# Dating-app

## What is it ?
It is an app inspired by tinder/bumble. This app have basic functionality such as:
* Sign up/Login
* User able to view, like/pass candidates with limitation of 10 unique dating profile daily
* User able to purchase premium packages such as:
    * No swipe quota 
    * Verified label for user

# Code Structure
Code file structure inspired by Clean Architecture, separation of concern between infrastructure, application logic and business 

### Current file structure:

```
├───application
│   ├───core
│   ├───dto
│   ├───usecase
│   └───utils
│       ├───date
│       ├───otp
│       └───phoneNumber
├───infrastructure
│   ├───db
│   │   └───inmemory
│   └───httpserver
│       ├───controller
│       │   └───models
│       ├───jwt
│       └───middlewares
└───mocks
```
`application/core` folder contains file concerning core business logic. `application/usecase` contains files concerning application use case. `infrastructure/` contain component, codes which deals interraction with outside services. 

### How to Run Service
1. Clone repo
2. Run `go mod tidy` to download dependencies
3. Run using `go run ./main.go`

### Development
In order to contribute, you might want to install some tools:
* `pre-commit` - `go install github.com/lietu/go-pre-commit@latest`
* `golangci-lint` - `go install github.com/golangci/golangci-lint/cmd/golangci-lint`
* `mockery` - `go install github.com/vektra/mockery/v2@latest`

# Homeworks:
There are some leftover task which haven't been implemented but important in order to improve system reliability, observability, maintainability
* Error object
* Logging
* More complete test
* Use of environment variables
* Use of dependency injection
* Use real database
* Docker build
* GIT commit,code check
# VadeTestApi
 
## init go module
    go mod init vade.com/vade/goDocumentAPI

## Installing gorilla/mux
    go get -u github.com/gorilla/mux

## run project
    go run main.go

## install goconvey(https://github.com/smartystreets/goconvey)
    go get github.com/smartystreets/goconvey
    go install github.com/smartystreets/goconvey
    go get github.com/smartystreets/goconvey/convey@v1.8.1

## Start up the GoConvey web server at your project's path:
    $ $GOPATH/bin/goconvey

## test results in the browser at
    http://localhost:8080

## Building Docker Image
    docker build -t microservice .

## run docker
    docker run -p 9090:9090 microservice
FROM golang:lastest

RUN go get github.com/labstack/echo \
    && go get github.com/sirupsen/logrus

ADD . $GOPATH/src/github.com/nolleh/goecho

RUN go test github.com/nolleh/goecho/...

WORKDIR $GOPATH/src/github.com/nolleh/goecho

CMD go run main.go
FROM  golang:1.19-alpine3.16  AS build-env
ENV ExpectedMode=FullMode
RUN echo $GOPATH
RUN apk add --no-cache git gcc musl-dev
COPY . /go/src/github.com/devtron-labs/automation-suite/
WORKDIR /go/src/github.com/devtron-labs/automation-suite/
ADD . /go/src/github.com/devtron-labs/automation-suite/
CMD [ "/bin/sh" , "-c" , "cd $ExpectedMode && go test -timeout 99999s . -v" ]
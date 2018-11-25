FROM golang:1.11-alpine3.7

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
ENV APPDIR $GOPATH/src/github.com/higasgt/gocon-lt/app

WORKDIR ${APPDIR}
COPY . ${APPDIR}

RUN apk add --no-cache bash ca-certificates curl gcc git musl-dev openssh tzdata

RUN curl -L -s https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 \ 
    -o /go/bin/dep
RUN chmod +x /go/bin/dep

RUN go get github.com/oxequa/realize

ENTRYPOINT ["./docker-entrypoint-local.sh"]
CMD ["realize", "start"]

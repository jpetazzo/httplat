FROM golang:alpine
RUN apk add git
WORKDIR /go/src/httplat
COPY main.go .
RUN go get
RUN go install
FROM alpine
COPY --from=0 /go/bin/httplat /usr/local/bin
EXPOSE 9080

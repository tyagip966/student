FROM golang:1.14 AS build
RUN mkdir student
ADD . /student
WORKDIR /student
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app ./cmd/main.go

EXPOSE 8080
ENTRYPOINT [ "app"]

FROM golang:1.10.0-alpine3.7

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
EXPOSE 8080

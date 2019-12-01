FROM golang:1.13-alpine

RUN apk --no-cache add git
RUN apk --no-cache add build-base

WORKDIR /go/src/shortener
COPY ./shortener/ .

RUN go get -d -v -t ./...
RUN go build

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
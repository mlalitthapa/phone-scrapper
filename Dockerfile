FROM golang:alpine AS builder

# install git
RUN apk add --update --no-cache git

# install dep
RUN go get -v github.com/golang/dep/cmd/dep

# setup working directory
RUN mkdir -p /go/src/github.com/mlalitthapa/phone-scrapper
ADD ./ /go/src/github.com/mlalitthapa/phone-scrapper
WORKDIR /go/src/github.com/mlalitthapa/phone-scrapper

COPY Gopkg.toml Gopkg.lock ./
# copies the Gopkg.toml and Gopkg.lock to WORKDIR

RUN dep ensure -v -vendor-only
# install the dependencies without checking for go code

# build and run the binary file
RUN go build -o main .
CMD ["./main"]
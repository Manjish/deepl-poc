FROM golang:1.20.2-alpine

# Required because go requires gcc to build
RUN apk add build-base git inotify-tools
RUN echo $GOPATH
RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY . /deepl_poc
WORKDIR /deepl_poc
RUN go mod download

CMD sh /deepl_poc/docker/run.sh

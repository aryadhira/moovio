FROM golang:alpine

RUN apk update && apk add git

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY libs libs
COPY services/backend/streamer services/backend/streamer

RUN go mod tidy
RUN go build -o streamer ./services/backend/streamer

ENTRYPOINT [ "./streamer" ]
FROM golang:alpine

RUN apk update && apk add git

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY libs libs
COPY services/backend/movies services/backend/movies

RUN go mod tidy
RUN go build -o movies ./services/backend/movies

ENTRYPOINT [ "./movies" ]
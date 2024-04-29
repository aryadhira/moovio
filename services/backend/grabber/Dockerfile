FROM golang:alpine

RUN apk update && apk add git

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY libs libs
COPY services/backend/grabber services/backend/grabber

RUN go mod tidy
RUN go build -o grabber ./services/backend/grabber

ENTRYPOINT [ "./grabber" ]
FROM golang:alpine

RUN apk update && apk add git

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY services/backend/api-gateway .

RUN go mod tidy
RUN go build -o apigateway

ENTRYPOINT [ "./apigateway" ]
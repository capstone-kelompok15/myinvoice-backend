# Build Stage
FROM golang:1.18-alpine as build

RUN apk update && apk upgrade

WORKDIR /app

COPY go.mod /app/
COPY go.sum /app/

RUN go mod download
RUN go mod tidy

COPY . /app/

RUN go build -o /app/main

# Execute Stage
FROM alpine:3.16.0

RUN apt-get update -y && apt-get upgrade -y
RUN apt-get install -y xvfb libfontconfig wkhtmltopdf

RUN apt-get install -y ca-certificates

WORKDIR /app

COPY ./assets/ ./assets/
COPY --from=build /app/main /app/main

EXPOSE 8080

CMD ["./main"]
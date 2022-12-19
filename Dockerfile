# Build Stage
FROM golang:1.18-bullseye as build

RUN apt-get update -y && apt-get upgrade -y

WORKDIR /app

COPY go.mod /app/
COPY go.sum /app/

RUN go mod download
RUN go mod tidy

COPY . /app/

RUN go build -o /app/main

# Execute Stage
FROM ubuntu:22.04

RUN apt-get update -y && apt-get upgrade -y
RUN apt-get install -y libfontconfig wkhtmltopdf
RUN apt-get install -y ca-certificates

ENV TZ=Asia/Makassar
RUN apt-get install -y tzdata
RUN cp /usr/share/zoneinfo/$TZ /etc/localtime

WORKDIR /app

COPY ./assets/ ./assets/
COPY --from=build /app/main /app/main

EXPOSE 8080

CMD ["./main"]
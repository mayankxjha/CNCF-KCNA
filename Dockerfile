FROM golang:tip-bookworm

COPY . /App

WORKDIR /App

EXPOSE 8080

CMD ["/bin/bash", "-c", "go run main.go"]

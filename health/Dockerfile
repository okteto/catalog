FROM golang:buster as builder

WORKDIR /
ADD . .
RUN go build ./cmd/main.go

EXPOSE 8080
CMD ["/main"]

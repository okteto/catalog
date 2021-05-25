FROM golang:buster as builder

WORKDIR /
ADD . .
RUN go build ./main.go

EXPOSE 8080
CMD ["/main"]

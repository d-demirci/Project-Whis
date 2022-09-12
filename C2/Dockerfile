# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /C2

COPY . .
RUN go mod init github.com/d-demirci/Project-Whis/C2
RUN go mod tidy
RUN go get github.com/roerohan/wait-for-it
RUN go install github.com/roerohan/wait-for-it
RUN go build -o /C2/bin/C2

EXPOSE 8080 5555 53899

CMD [ "/C2/bin/C2" ]
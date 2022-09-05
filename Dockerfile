FROM golang:1.17-buster as builder

ENV APP_HOME /verify

RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"
COPY . "$APP_HOME"

RUN go mod download && go mod verify

RUN go build app/main.go

CMD ["./run.sh"]
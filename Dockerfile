FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY  *.go ./

RUN go build -o /goDocker

EXPOSE 4000

CMD ["/goDocker"]
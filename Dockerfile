FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY * ./

RUN go build -o /go-docker-resize

EXPOSE 8080

CMD [ "/go-docker-resize" ]
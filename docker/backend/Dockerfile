FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY database ./

RUN go mod download

COPY . .

RUN go build -o main ./src/cmd

CMD ["./main"]
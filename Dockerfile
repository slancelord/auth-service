FROM golang:1.24.0

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o auth-service ./cmd

EXPOSE 8080

CMD ["./auth-service"]
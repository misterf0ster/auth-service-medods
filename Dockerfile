FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd

RUN go build -o auth-app

CMD ["/app/cmd/auth-app"]
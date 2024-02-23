FROM golang:1.21-bullseye

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /todo-go

EXPOSE 8080

CMD ["/todo-go"]
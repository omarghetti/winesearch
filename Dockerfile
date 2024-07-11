FROM golang:1.22-bullseye

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN GCO_ENABLED=0 GOOS=linux go build -o /winesearch

EXPOSE 8080

CMD ["/winesearch"]
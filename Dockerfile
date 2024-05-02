FROM golang:1.18 as builder

WORKDIR /app

COPY ../backend/go.mod ../backend/go.sum ./
RUN go mod download

COPY ../backend .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o backend_app ./cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/backend_app .

COPY ../.env ./

EXPOSE 3001

CMD ["./backend_app"]

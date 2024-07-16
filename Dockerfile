FROM golang:1.22 AS builder

WORKDIR /app

COPY . .

RUN go get -d -v ./...
RUN GOOS=linux GOARCH=arm go build -o app cmd/main.go

FROM golang:1.22-alpine

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app
RUN chown -R appuser:appgroup /app

USER appuser

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["/app/app"]

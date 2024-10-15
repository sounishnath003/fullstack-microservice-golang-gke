FROM golang:1.23 AS builder

WORKDIR /app

COPY . .
RUN go mod tidy
RUN go mod download -x
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /app/bin/goblogsservice main.go


FROM scratch
WORKDIR /app

COPY queries.sql .
COPY --from=builder /app/bin/goblogsservice .

EXPOSE 3001

CMD ["/app/goblogsservice"]

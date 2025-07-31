# syntax=docker/dockerfile:1
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM gcr.io/distroless/static:nonroot

WORKDIR /

COPY --from=builder /app/server /server

EXPOSE 8080

USER nonroot:nonroot
ENTRYPOINT ["/server"]

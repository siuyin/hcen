FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN mkdir -p bin && CGO_ENABLED=0 go build -o bin -v ./cmd/hcen-pos && ls -lF bin

FROM scratch
COPY --from=builder /app/bin/hcen-pos /app
EXPOSE 8080
CMD ["/app"]

# vim: set filetype=dockerfile :

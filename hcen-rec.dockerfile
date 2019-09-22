FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN mkdir -p bin && CGO_ENABLED=0 go build -o bin -v ./cmd/hcen-rec && ls -lF bin

FROM scratch
COPY --from=builder /app/bin/hcen-rec /app
EXPOSE 8080
CMD ["/app"]

# vim: set filetype=dockerfile :

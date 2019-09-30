FROM golang:1.13.1 AS builder
RUN apt update && apt -yqq install libxml2
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && mkdir -p /db2 && go run /go/pkg/mod/github.com/ibmdb/go_ibm_db@v0.1.1/installer /db2
COPY . .
ENV DB2HOME=/db2/clidriver CGO_CFLAGS=-I/db2/clidriver/include CGO_LDFLAGS=-L/db2/clidriver/lib LD_LIBRARY_PATH=/db2/clidriver/lib
RUN mkdir -p bin && CGO_ENABLED=1 go build -o bin -v ./cmd/hcen-rec

FROM ubuntu:18.04
RUN apt update && apt -yqq install libxml2
COPY --from=builder /db2/ /db2
COPY --from=builder /app/bin/hcen-rec /app
ENV DB2HOME=/db2/clidriver CGO_CFLAGS=-I/db2/clidriver/include CGO_LDFLAGS=-L/db2/clidriver/lib LD_LIBRARY_PATH=/db2/clidriver/lib
EXPOSE 8080
CMD ["/app"]

# vim: set filetype=dockerfile :

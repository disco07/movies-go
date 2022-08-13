FROM golang:1.18 AS builder
COPY go.mod go.sum /go/src/github.com/disco07/movies-go/
WORKDIR /go/src/github.com/disco07/movies-go/
RUN go mod download
COPY . /go/src/github.com/disco07/movies-go/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/movies-go github.com/disco07/movies-go

FROM scratch
COPY --from=builder /go/src/github.com/disco07/movies-go/ /usr/bin/movies-go
EXPOSE 4001 4001
ENTRYPOINT ["/usr/bin/movies-go"]
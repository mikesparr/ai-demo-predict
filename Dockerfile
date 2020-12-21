FROM golang:1.15.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/github.com/mikesparr/ai-demo-predict/
WORKDIR /go/src/github.com/mikesparr/ai-demo-predict
RUN go mod download
COPY . /go/src/github.com/mikesparr/ai-demo-predict
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/ai-demo-predict github.com/mikesparr/ai-demo-predict

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/mikesparr/ai-demo-predict/build/ai-demo-predict /usr/bin/ai-demo-predict
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/ai-demo-predict"]
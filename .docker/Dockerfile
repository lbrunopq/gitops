FROM golang:1.15
WORKDIR /go/src/k8s
COPY ./src .
RUN GOOS=linux go build -ldflags="-s -w"
CMD ["./k8s"]
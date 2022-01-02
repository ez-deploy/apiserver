FROM golang:1.17 as builder

WORKDIR /go/src/github.com/ez-deploy/apiserver
COPY . .

RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.io && \
    go mod tidy && \
    CGO_ENABLED=0 go build -tags netgo -o apiserver ./cmd/ez-deploy-apiserver-server/main.go

FROM busybox

WORKDIR /

COPY --from=builder /go/src/github.com/ez-deploy/apiserver/apiserver /apiserver

EXPOSE 80
ENTRYPOINT [ "/apiserver", "--port=80", "--host=\"0.0.0.0\"" ]
FROM nightlord189/golangweb:latest AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM scratch

COPY --from=builder /build/main /
COPY --from=builder /build/config.json /

EXPOSE 8080

ENTRYPOINT ["/main"]
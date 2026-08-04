# dev
FROM golang:1.26-bookworm AS dev
WORKDIR /work/yatter-backend-go
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.1
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.1.6

COPY ./ ./
RUN make mod build-linux

# release
FROM alpine AS release
RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
COPY --from=dev /work/yatter-backend-go/build/yatter-backend-go-linux-amd64 /usr/local/bin/yatter-backend-go
EXPOSE 8080
ENTRYPOINT ["yatter-backend-go"]

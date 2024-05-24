FROM golang:1.22.2

RUN mkdir -p /app/atcoder-web-app

WORKDIR /app/atcoder-web-app

COPY . .

RUN go install github.com/cespare/reflex@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

CMD reflex -r '(\.go$|go\.mod)' -s go run *.go
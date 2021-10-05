FROM golang
ENV GO111MODULE on
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /app
COPY . .
RUN go build ./cmd/main.go
ENTRYPOINT ./main
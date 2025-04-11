FROM golang:1.24-alpine AS builder
ADD . /go/src/github.com/pakohan/cooklang-epub
WORKDIR /go/src/github.com/pakohan/cooklang-epub
RUN go build -o /go/bin/cooklang-epub .

FROM alpine:latest
RUN apk add --no-cache tzdata
COPY --from=builder /go/bin/cooklang-epub /usr/local/bin/cooklang-epub
CMD [ "/usr/local/bin/cooklang-epub" ]

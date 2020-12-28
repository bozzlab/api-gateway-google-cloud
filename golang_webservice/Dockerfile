FROM golang:1.13.4-alpine as builder
RUN adduser -D -g '' appuser
RUN apk update && apk add gcc git
WORKDIR /opendx-cloud-build
COPY go.mod .
COPY . .
RUN go build ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /opendx-cloud-build/ .
CMD ["./app"] 
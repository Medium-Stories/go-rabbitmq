# build app
FROM golang:1.19-alpine3.16 as base_build

RUN apk add gcc libc-dev libffi-dev

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .
RUN go build -v -o order-svc cmd/order/main.go

# create runtime
FROM alpine:3.15.0

WORKDIR /app

RUN apk add ca-certificates

COPY --from=base_build /app/order-svc .

EXPOSE 8001

ENTRYPOINT ["/app/order-svc"]
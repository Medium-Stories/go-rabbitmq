# build app
FROM golang:1.19-alpine3.16 as base_build

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .
RUN go build -v -o shipping-svc cmd/shipping/main.go

# create runtime
FROM alpine:3.15.0

WORKDIR /app

RUN apk add ca-certificates

COPY --from=base_build /app/shipping-svc .

EXPOSE 8003

ENTRYPOINT ["/app/shipping-svc"]
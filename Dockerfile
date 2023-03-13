FROM golang:1.19.3-alpine as builder
WORKDIR /build
COPY go.mod . 
COPY go.sum . 

RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /jwt_registration_api cmd/main.go

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /bin/backend

COPY --from=builder /build/config ./config
COPY --from=builder jwt_registration_api ./jwt_registration_api

ENTRYPOINT ["./jwt_registration_api"]

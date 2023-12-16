FROM golang:1.21.5 as buider

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o ingress-manager main.go

FROM alpine:latest

WORKDIR /app

COPY --from=buider /app/ingress-manager .

CMD ["./ingress-manager"]
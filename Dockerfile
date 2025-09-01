FROM golang:1-alpine as builder
WORKDIR /
COPY . ./
RUN go mod download

RUN go build -o /health-server

FROM alpine
COPY --from=builder /health-server .

EXPOSE 80
CMD [ "/health-server" ]

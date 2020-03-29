FROM golang:1.14 as builder

WORKDIR /go/src/ra
COPY . .

RUN CGO_ENABLED=0 go build -o /bin/ra


FROM alpine

WORKDIR /bin/ra

COPY --from=builder /bin/ra .
COPY roles.json .

CMD ["./ra"]

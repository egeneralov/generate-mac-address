FROM golang:alpine

WORKDIR /go/src/github.com/egeneralov/generate_mac_address/

ADD . .
RUN go build -o /go/bin/generate_mac_address main.go

FROM alpine
COPY --from=0 /go/bin/generate_mac_address /
ENV PORT=8080
ENTRYPOINT ["/generate_mac_address"]

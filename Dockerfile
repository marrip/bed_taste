FROM golang:1.16.4-buster AS builder
COPY . /bed_taste
WORKDIR /bed_taste
RUN go build .

FROM debian:buster-20210511-slim
COPY --from=builder /bed_taste/bed_taste /usr/local/bin/

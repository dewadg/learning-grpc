FROM golang:1.14-alpine AS build

WORKDIR /build

ADD . .

RUN go get -v
RUN go build -o api

FROM alpine

WORKDIR /usr/local/bin

COPY --from=build /build/api .

CMD ["api", "http:serve"]

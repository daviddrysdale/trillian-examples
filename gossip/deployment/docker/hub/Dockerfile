FROM golang:1.10 as build

ADD . /go/src/github.com/google/trillian-examples
WORKDIR /go/src/github.com/google/trillian-examples

ARG GOFLAGS=""
RUN go get ${GOFLAGS} ./gossip/hub/hub_server

FROM gcr.io/distroless/base

COPY --from=build /go/bin/hub_server /

ENTRYPOINT ["/hub_server"]

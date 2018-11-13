FROM golang:1.11 AS build

ARG REPOSITORY="github.com/justinm/tfjson2"
ENV GOPATH=/build
WORKDIR /build/src/$REPOSITORY/

COPY vendor/ /build/src/$REPOSITORY/vendor/
COPY tfjson2.go Gopkg.lock Gopkg.toml /build/src/$REPOSITORY/
COPY tfjson2/ /build/src/$REPOSITORY/tfjson2/

RUN go build -o /build/bin/tfjson2

FROM golang:1.11

COPY --from=build /build/bin/tfjson2 /usr/local/bin/tfjson2

ENTRYPOINT ["/usr/local/bin/tfjson2"]

CMD ["-h"]

FROM golang:1.11 AS build

ARG TFJSON_REPOSITORY="github.com/cloudvar/tfjson_with_open-policy-agent"
ENV GOPATH=/build
WORKDIR /build/src/$TFJSON_REPOSITORY/

COPY vendor/ /build/src/$TFJSON_REPOSITORY/vendor/
COPY tfjson2.go Gopkg.lock Gopkg.toml /build/src/$TFJSON_REPOSITORY/
COPY tfjson2/ /build/src/$TFJSON_REPOSITORY/tfjson2/

RUN go build -o /build/bin/tfjson2

FROM golang:1.11

COPY --from=build /build/bin/tfjson2 /usr/local/bin/tfjson2

# Adding OPA
RUN curl -L -o /opa https://github.com/open-policy-agent/opa/releases/download/v0.10.1/opa_linux_amd64
RUN chmod 755 /opa

# Add Policies
COPY policies /opt/policies

# ENTRYPOINT ["/usr/local/bin/tfjson2"]

# CMD ["-h"]

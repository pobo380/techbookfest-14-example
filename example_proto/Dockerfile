FROM golang:1.19

ARG PROTOBUF_VERSION=22.5

# Install unzip command.
RUN apt-get update && apt-get -y install unzip curl

# Install protoc command.
WORKDIR /tmp/protoc
RUN curl -L https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOBUF_VERSION}/protoc-${PROTOBUF_VERSION}-linux-x86_64.zip -o protoc.zip && \
    unzip protoc.zip && \
    mv bin/* /usr/local/bin/ && \
    mv include/* /usr/local/include/

# Install protoc-gen-go
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30

ENTRYPOINT ["protoc"]

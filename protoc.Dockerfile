FROM darthcoder4309/protoc-dart

ARG APT_INSTALL="apt-get install -y --no-install-recommends"

RUN DEBIAN_FRONTEND=noninteractive apt-get update && \
    ${APT_INSTALL} \
        golang

RUN go get google.golang.org/protobuf/cmd/protoc-gen-go

ENV GOPATH=/root/go
ENV PATH=$PATH:$GOPATH/bin

COPY ./scripts/compile_protobuf.sh /compile.sh
RUN chmod +x /compile.sh

ENTRYPOINT [ "/compile.sh" ]

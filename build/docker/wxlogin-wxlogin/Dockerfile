ARG REGISTRY
#  builder
FROM ${REGISTRY}/golang:1.23 as builder

WORKDIR /opt/build
COPY wxlogin wxlogin
RUN cd wxlogin && go build  -o wxlogin -v ./cmd/wxlogin

FROM ${REGISTRY}/ubuntu:24.04

RUN apt-get update && \
    apt-get install -y ca-certificates tzdata && \
    rm -rf /var/lib/apt/lists/* && \
    ln -fs /usr/share/zoneinfo/Etc/UTC /etc/localtime && \
    dpkg-reconfigure --frontend noninteractive tzdata

COPY --from=builder /opt/build/wxlogin /opt/wxlogin/

WORKDIR /opt/wxlogin
ENTRYPOINT ["/opt/wxlogin/wxlogin"]

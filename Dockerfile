FROM golang:1.13.6

ENV GRPCURL_VERSION=1.4.0
ENV GRPCURL_NAME=grpcurl_${GRPCURL_VERSION}_linux_x86_64
ENV GRPCURL_FILENAME=${GRPCURL_NAME}.tar.gz
ENV GRPCURL_URL=https://github.com/fullstorydev/grpcurl/releases/download/v${GRPCURL_VERSION}/${GRPCURL_FILENAME}
RUN cd /tmp \
    && curl -sL ${GRPCURL_URL} | tar xz \
    && mv /tmp/grpcurl /usr/local/bin \
    && ln /usr/local/bin/grpcurl /usr/local/bin/gurl \
    && rm -rf /tmp/*

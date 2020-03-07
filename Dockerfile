FROM golang:1.13.6

# install system deps
RUN apt-get update \
    && apt-get install unzip

# fetch/install protobuf compiler
ENV PROTOC_VERSION 3.11.4
ENV PROTOC_PATH protoc-${PROTOC_VERSION}-linux-x86_64
ENV PROTOC_FILE ${PROTOC_PATH}.zip
ENV PROTOC_URL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/${PROTOC_FILE}
RUN mkdir -p /opt/src \
    && curl -LO ${PROTOC_URL} \
    && unzip ${PROTOC_FILE} -d /opt/src/${PROTOC_PATH} \
    && ln -s /opt/src/${PROTOC_PATH}/bin/protoc /usr/local/bin/ \
    && cp --recursive /opt/src/${PROTOC_PATH}/include/google /usr/local/include \
    && rm ${PROTOC_FILE}

# fetch/install grpcurl
ENV GRPCURL_VERSION=1.4.0
ENV GRPCURL_NAME=grpcurl_${GRPCURL_VERSION}_linux_x86_64
ENV GRPCURL_FILENAME=${GRPCURL_NAME}.tar.gz
ENV GRPCURL_URL=https://github.com/fullstorydev/grpcurl/releases/download/v${GRPCURL_VERSION}/${GRPCURL_FILENAME}
RUN cd /tmp \
    && curl -sL ${GRPCURL_URL} | tar xz \
    && mv /tmp/grpcurl /usr/local/bin \
    && ln /usr/local/bin/grpcurl /usr/local/bin/gurl \
    && rm -rf /tmp/*

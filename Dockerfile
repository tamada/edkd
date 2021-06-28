FROM alpine:3.10.1

ARG version=1.0.0

LABEL maintainer="Katsuyuki Yamauchi" \
      description="Similarities and distances calculator among vectors"

RUN    adduser -D edkd \
    && apk --no-cache add --update --virtual .builddeps curl tar \
#    && curl -s -L -O https://github.com/YKatsuy/edkd/releases/download/v${version}/edkd-${version}_linux_amd64.tar.gz \
    && curl -s -L -o edkd-${version}_linux_amd64.tar.gz https://drive.google.com/drive/folders/1KgmYWBnHmt-SrMmrFMp-W2BBZv-1c2gj \
    && tar xfz edkd-${version}_linux_amd64.tar.gz        \
    && mv edkd-${version} /opt                           \
    && ln -s /opt/edkd-${version} /opt/edkd               \
    && ln -s /opt/edkd-${version}/edkd /usr/local/bin/edkd \
    && rm edkd-${version}_linux_amd64.tar.gz             \
    && apk del --purge .builddeps

ENV HOME="/home/edkd"

WORKDIR /home/edkd
USER    edkd

ENTRYPOINT [ "/opt/edkd/edkd" ]
FROM ARG_FROM

# Build-time metadata as defined at http://label-schema.org
ARG BUILD_DATE
ARG VCS_REF
ARG VERSION
ARG DOCKER_TAG

LABEL org.label-schema.build-date=$BUILD_DATE \
      org.label-schema.name="Currency Exchange Server" \
      org.label-schema.description="Currency Exchange Server" \
      org.label-schema.url="https://github.com/me-io/go-swap" \
      org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.vcs-url="https://github.com/me-io/go-swap" \
      org.label-schema.vendor="ME.IO" \
      org.label-schema.version=$VERSION \
      org.label-schema.schema-version="$DOCKER_TAG"

RUN apk update \
    && apk upgrade \
    && apk add --no-cache ca-certificates \
    && update-ca-certificates

ADD bin/ARG_OS-ARG_ARCH/ARG_SRC_BIN /ARG_BIN
ENV BINSRC_ENV="/ARG_BIN"
COPY scripts/docker-entrypoint.sh /usr/local/bin/

EXPOSE 5000

USER nobody:nobody
ENTRYPOINT ["docker-entrypoint.sh"]

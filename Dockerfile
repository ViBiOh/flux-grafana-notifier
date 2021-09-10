FROM vibioh/scratch

ENV API_PORT 1080
EXPOSE 1080

COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

HEALTHCHECK --retries=5 CMD [ "/notifier", "-url", "http://localhost:1080/health" ]
ENTRYPOINT [ "/notifier" ]

ARG VERSION
ENV VERSION=${VERSION}

ARG TARGETOS
ARG TARGETARCH

COPY release/notifier_${TARGETOS}_${TARGETARCH} /notifier

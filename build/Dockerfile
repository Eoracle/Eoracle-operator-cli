# Specify arguments for flexibility
ARG GO_VERSION=1.23-alpine
ARG BASE=golang:${GO_VERSION}
ARG BUILDPLATFORM=linux

# Stage 1: Builder
FROM --platform=${BUILDPLATFORM} ${BASE} AS builder

# Install dependencies
RUN apk update && apk add --no-cache \
    linux-headers \
    gcc \
    libtool \
    openssl-dev \
    libffi \
    tini \
    git \
    'su-exec>=0.2' \
    build-base \
    libffi-dev \
    openssh-client

# Add user and group
RUN addgroup -g 502 -S eoracle \
    && adduser -u 502 -S -G eoracle eoracle \
    && mkdir -p /app/eoracle/logs \
    && chown -R eoracle:eoracle /app/eoracle

# Set the working directory
WORKDIR /app/eoracle

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy application code
COPY --chown=eoracle:eoracle . .

# Build the binary
RUN platform=$(uname -m) && \
    [ "$platform" = "x86_64" ] && platform=amd64 || platform=arm64 && \
    make build-oprcli-${platform} && \
    mv ./bin/oprcli-${platform} ./bin/oprcli && \
    chown eoracle:eoracle ./bin/oprcli

FROM --platform=$BUILDPLATFORM ${BASE}

# Install runtime dependencies
RUN apk update \
 && apk add --no-cache \
    linux-headers \
    gcc \
    libtool \
    openssl-dev \
    libffi \
    tini \
    git \
    'su-exec>=0.2' \
 && apk add --no-cache --virtual .build_deps build-base libffi-dev


# Add user and group
RUN addgroup -g 502 -S eoracle \
    && adduser -u 502 -S -G eoracle eoracle \
    && mkdir -p /app/eoracle/ \
    && chown -R eoracle:eoracle /app/eoracle

# Set the working directory
WORKDIR /app/eoracle

# Copy binary and entrypoint script
COPY --from=builder /app/eoracle/bin/oprcli /usr/local/bin/oprcli

COPY build/docker-entrypoint.sh /usr/local/bin/

ENV PATH="/app/eoracle/bin:${PATH}"

ENTRYPOINT ["/sbin/tini", "--", "/bin/sh", "/usr/local/bin/docker-entrypoint.sh"]

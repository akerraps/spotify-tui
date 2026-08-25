# ─────────────────────────────────────────────
# Stage 1: Clone & build TuneCTL
# ─────────────────────────────────────────────

FROM golang:1.26-bookworm AS builder

ARG TUNECTL_REPO=https://github.com/akerraps/tunectl.git
ARG TUNECTL_VERSION=main

RUN apt-get update \
    && apt-get install -y --no-install-recommends git \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /src

RUN git clone --branch ${TUNECTL_VERSION} --depth 1 \
    ${TUNECTL_REPO} .

RUN CGO_ENABLED=0 go build \
    -o /tunectl \
    ./cmd/tunectl


# ─────────────────────────────────────────────
# Stage 2: Runtime
# ─────────────────────────────────────────────

FROM debian:bookworm-slim

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        ca-certificates \
        curl \
        unzip \
        python3 \
        ffmpeg \
        nodejs \
    && rm -rf /var/lib/apt/lists/*

# Create non-root user
RUN useradd \
    --create-home \
    --shell /bin/bash \
    tunectl

# Copy compiled binary
COPY --from=builder /tunectl /usr/local/bin/tunectl

# Run everything as non-root
USER tunectl

ENV HOME="/home/tunectl"
ENV PATH="/home/tunectl/.deno/bin:${PATH}"

# Install Deno
RUN curl -fsSL https://deno.land/install.sh | sh

WORKDIR /music

ENTRYPOINT ["tunectl"]
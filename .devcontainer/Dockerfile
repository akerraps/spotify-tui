FROM mcr.microsoft.com/devcontainers/go:1.26-bookworm

USER root

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        gcc \
        python3 \
        ffmpeg \
        nodejs \
    && rm -rf /var/lib/apt/lists/*
    
RUN curl -fsSL https://deno.land/install.sh | sh \
    && mv /root/.deno /home/vscode/.deno \
    && chown -R vscode:vscode /home/vscode/.deno
    
    ENV PATH="/home/vscode/.deno/bin:${PATH}"

RUN chown -R vscode:vscode /go

RUN go install golang.org/x/tools/gopls@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest
    
WORKDIR /workspaces/tunectl

COPY --chown=vscode:vscode go.mod go.sum ./

RUN go mod download

USER vscode
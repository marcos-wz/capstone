FROM golang:1.17.6-buster 

LABEL org.opencontainers.image.authors="Marcos Jauregui <marcos.jauregui@wizeline.com>" \
	org.opencontainers.image.title="CAPSTONE-DEV" \
    org.opencontainers.image.description="Wizeline GO Bootcamp - Capstone Project - Development" 

# CREATE DEFAULT USER ───────────────────────────────────────────────────────

ARG USER="capstone" 
RUN set -ex \
    && useradd \
        --comment "capstone user" \
        --create-home \
        --no-log-init \
        --home-dir /home/${USER} \
        ${USER} 

# USER SCOPE ────────────────────────────────────────────────────────

USER ${USER}

RUN set -ex \
    # HTTP Server
    && go get github.com/labstack/echo/v4 \
    # TESTING
    && go get github.com/stretchr/testify \
    # VSCODE GO Extensions
	&& go install golang.org/x/tools/gopls@latest \
    && go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest \
    && go install github.com/ramya-rao-a/go-outline@latest \
    && go install honnef.co/go/tools/cmd/staticcheck@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    # DEVELOPMENT TOOLS
    && go get github.com/githubnemo/CompileDaemon

# RUN ────────────────────────────────────────────────────────

EXPOSE 8080

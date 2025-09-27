FROM mcr.microsoft.com/devcontainers/go:1.25

# Install GoNB kernel for Jupyter
RUN go install github.com/janpfeifer/gonb@latest
RUN go install golang.org/x/tools/cmd/goimports@latest
RUN go install golang.org/x/tools/gopls@latest
RUN gonb --install

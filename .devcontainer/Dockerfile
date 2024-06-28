FROM golang:1.21

# especificação de usuário para contexto mapeamento volumes
ARG USERNAME=gopher
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && chown -R $USER_UID:$USER_GID /home/$USERNAME

USER ${USERNAME}
ENV GOPATH="/home/${USERNAME}"
ENV PATH="${GOPATH}/bin:${PATH}"

RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/tpng/gopkgs@latest
RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install github.com/cosmtrek/air@latest

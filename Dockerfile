FROM ubuntu:24.04

# Atualize os pacotes do sistema
RUN apt-get update && apt-get upgrade -y

# Instale dependências essenciais
RUN apt-get install -y \
    curl \
    unzip \
    git \
    build-essential \
    software-properties-common \
    zsh \
    wget \
    locales

# Configure o locale para UTF-8 (importante para evitar problemas com caracteres)
RUN locale-gen en_US.UTF-8 && update-locale LANG=en_US.UTF-8

# Adicione Go (ajuste a versão conforme necessário, aqui está a versão 1.21 como exemplo)
ENV GO_VERSION=1.23.5
RUN curl -OL https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz \
    && rm go${GO_VERSION}.linux-amd64.tar.gz

# Configure o PATH para incluir o binário do Go
ENV PATH="/usr/local/go/bin:$PATH"

# Instale o oh-my-zsh
RUN sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" "" --unattended

# Configure um tema e plugins padrão para o Zsh
RUN git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions \
    && git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting

# Configure o Zsh como shell padrão
RUN chsh -s $(which zsh)

# Instale extensões de suporte ao Golang
RUN go install golang.org/x/tools/gopls@latest

# Configure um workspace Go
ENV GOPATH=/go
RUN mkdir -p /go/bin /go/src/workspace
WORKDIR /go/src/workspace

# Finalize limpando pacotes desnecessários
RUN apt-get autoremove -y && apt-get clean

# Exponha as portas (se for necessário para debug via DAP)
EXPOSE 8080

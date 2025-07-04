ARG IMAGE=ubuntu
ARG TAG=24.04
ARG MAKE_JOBS=2


FROM ${IMAGE}:${TAG} AS go-builder

RUN DEBIAN_FRONTEND=noninteractive apt-get update && \
    apt-get install -y -qq \
    curl \
    > /dev/null 2>&1 && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

ENV GO_VERSION=1.23.10
RUN curl --fail -s https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz | tar -C /usr/local -xz


FROM ${IMAGE}:${TAG} AS install

RUN DEBIAN_FRONTEND=noninteractive apt-get update && \
    apt-get install -y -qq \
    apt-transport-https \
    build-essential \
    ca-certificates \
    curl \
    git \
    gnupg-agent \
    jq \
    make \
    software-properties-common \
    sudo \
    zsh \
    > /dev/null 2>&1 && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

ARG TARGETARCH
RUN curl --fail -s https://github.com/mikefarah/yq/releases/latest/download/yq_linux_${TARGETARCH} -o /usr/local/bin/yq && chmod +x /usr/local/bin/yq

COPY --from=go-builder /usr/local/go /usr/local/go

ENV PATH=/usr/local/go/bin:$PATH

SHELL ["/usr/bin/zsh", "-o", "pipefail", "-c"]

CMD /usr/bin/zsh

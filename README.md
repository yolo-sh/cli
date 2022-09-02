<p align="center">
  <img src="https://user-images.githubusercontent.com/1233275/188123619-b37cd965-cdbf-4eb6-9218-fe81c136ed8a.jpeg" alt="yolo" width="180" height="180" />
</p>

<p align="center">
    <h1 align="center">Yolo</h1>
    <p align="center">Live environments for any repository. Running in your cloud provider account.<br />Currently available on <a href="https://github.com/yolo-sh/aws-cloud-provider">Amazon Web Services</a> and <a href="https://github.com/yolo-sh/hetzner-cloud-provider">Hetzner</a>.</p>
</p>

<blockquote align="left">
  ... you could use it to deploy your app, as a remote development environment or even as a Ngrok replacement. Honestly, it's up to you!
</blockquote>

```bash
yolo aws init yolo-sh/api --instance-type t2.medium
```

https://user-images.githubusercontent.com/1233275/172346442-d6fef09c-2ef0-4633-8d72-e20bef8fc1a9.mp4

<img width="1136" alt="vscode" src="https://user-images.githubusercontent.com/1233275/172015213-0ba516b6-fe24-4bd4-8ad4-d876c6188f3c.png">

## Table of contents
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
    - [Login](#login)
    - [Init](#init)
    - [Edit](#edit)
    - [Open port](#open-port)
    - [Close port](#close-port)
    - [Remove](#remove)
    - [Uninstall](#uninstall)
- [Environments configuration](#environments-configuration)
- [Frequently asked questions](#frequently-asked-questions)
    - [Why using Docker as a VM and not something like NixOS, for example?](#-why-using-docker-as-a-vm-and-not-something-like-nixos-for-example)
    - [Given that my environment will run in a container does it mean that it will be limited?](#-given-that-my-environment-will-run-in-a-container-does-it-mean-that-it-will-be-limited)
- [License](#license)

## Requirements

The Yolo binary has been tested on Linux and Mac. Support for Windows is theoretical ([testers needed](https://github.com/yolo-sh/cli/issues/1) 💙).

Before using Yolo, the following dependencies need to be installed:

- [Visual Studio Code](https://code.visualstudio.com/) (currently the sole editor supported).

- [OpenSSH Client](https://www.openssh.com/) (used to access your environments).

## Installation

The easiest way to install Yolo is by running the following command in your terminal:

```bash
curl -sf https://raw.githubusercontent.com/yolo-sh/cli/main/install.sh | sh -s -- -b /usr/local/bin latest
```

This command could be run as-is or by changing:

  - The installation directory by replacing `/usr/local/bin` with your **preferred path**.
  
  - The version installed by replacing `latest` with a **[specific version](https://github.com/yolo-sh/cli/releases)**.

Once done, you could confirm that Yolo is installed by running the `yolo` command:

```bash
yolo --help
```

## Usage

```console
Yolo - Live environments for any repository in any cloud provider

To begin, run the command "yolo login" to connect your GitHub account.	

From there, the most common workflow is:

  - yolo <cloud_provider> init <repository>   : to initialize an environment for a specific GitHub repository

  - yolo <cloud_provider> edit <repository>   : to connect your preferred editor to an environment

  - yolo <cloud_provider> remove <repository> : to remove an unused environment
	
<repository> may be relative to your personal GitHub account (eg: cli) or fully qualified (eg: my-organization/api).

Usage:
  yolo [command]

Available Commands:
  aws         Use Yolo on Amazon Web Services
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  hetzner     Use Yolo on Hetzner
  login       Connect a GitHub account to use with Yolo

Flags:
  -h, --help      help for yolo
  -v, --version   version for yolo

Use "yolo [command] --help" for more information about a command.
```

### Login

```bash
yolo login
```
To begin, you need to run the `login` command to connect your GitHub account.

Yolo requires the following permissions:

  - "*Public SSH keys*" and "*Repositories*" to let you access your repositories from your environments.
	
  - "*GPG Keys*" and "*Personal user data*" to configure Git and sign your commits (verified badge).

**All your data (including the OAuth access token) are only stored locally in `~/.config/yolo/yolo.yml` (or in `XDG_CONFIG_HOME` if set).**

The source code that implements the GitHub OAuth flow is located in the [yolo-sh/api](https://github.com/yolo-sh/api) repository.

### Init

```bash
yolo <cloud_provider> start <repository>
```
The `init` command initializes an environment for a specific GitHub repository.

An `--instance-type` flag could be passed to specify the instance type that will power your environment. (*See the corresponding cloud provider repository for default / valid values*).

#### Examples

```bash
yolo aws init yolo-sh/api
```

```bash
yolo hetzner init yolo-sh/api --instance-type cx11
```

### Edit

```bash
yolo <cloud_provider> stop <repository>
```
The `edit` command connects your preferred editor to an environment.

Currently, the sole editor supported is **Visual Studio Code**. 

#### Example

```bash
yolo aws edit yolo-sh/api
```

### Open port

```bash
yolo <cloud_provider> open-port <repository> <port>
```
The `open-port` makes the specified port reachable from any IP address using the `TCP` protocol.

#### Example

```bash
yolo aws open-port yolo-sh/api 8000
```

### Close port

```bash
yolo <cloud_provider> close-port <repository> <port>
```
The `close-port` makes the specified port unreachable from all IP addresses.

#### Example

```bash
yolo aws close-port yolo-sh/api 8000
```

### Remove

```bash
yolo <cloud_provider> remove <repository>
```

The `remove` command removes an existing environment.

Removing means that the underlying instance **<ins>and all your data</ins>** will be **<ins>permanently removed</ins>**.

#### Example

```bash
yolo aws remove yolo-sh/api
```

### Uninstall

```bash
yolo <cloud_provider> uninstall
```

The `uninstall` command removes all the infrastructure components used by Yolo from your cloud provider account. (*See the corresponding cloud provider repository for details*).

**Before running this command, all environments need to be removed.**

#### Example

```bash
yolo aws uninstall
```

## Environments configuration

All environments are run in Docker containers built from the `yolosh/base-env:latest` image. You could see the source of this Docker image in the [yolo-sh/base-env](https://github.com/yolo-sh/base-env) repository:

```Dockerfile
# All environments will be Ubuntu-based (Ubuntu 22.04)
FROM buildpack-deps:jammy

ARG DEBIAN_FRONTEND=noninteractive

# RUN will use bash
SHELL ["/bin/bash", "-c"]

# We want a "standard Ubuntu"
# (ie: not one that has been minimized
# by removing packages and content
# not required in a production system)
RUN yes | unminimize

# Install system dependencies
RUN set -euo pipefail \
  && apt-get --assume-yes --quiet --quiet update \
  && apt-get --assume-yes --quiet --quiet install \
    apt-transport-https \
    build-essential \
    ca-certificates \
    curl \
    git \
    gnupg \
    locales \
    lsb-release \
    man-db \
    manpages-posix \
    nano \
    software-properties-common \
    sudo \
    tzdata \
    unzip \
    vim \
    wget \
  && apt-get clean && rm --recursive --force /var/lib/apt/lists/* /tmp/*

# Force LibSSL to 1.1.1 to avoid conflicts 
# with old Ruby and Python versions
RUN set -euo pipefail \
  && apt-add-repository --yes ppa:rael-gc/rvm \
  && apt-get --assume-yes --quiet --quiet update \
  && apt-get --assume-yes --quiet --quiet remove libssl-dev \
  && touch /etc/apt/preferences.d/rael-gc-rvm-precise-pin-900 \
  && { echo 'Package: *'; \
    echo 'Pin: release o=LP-PPA-rael-gc-rvm'; \
    echo 'Pin-Priority: 900'; } >> /etc/apt/preferences.d/rael-gc-rvm-precise-pin-900 \
  && apt-get --assume-yes --quiet --quiet install libssl-dev \
  && apt-get clean && rm --recursive --force /var/lib/apt/lists/* /tmp/*

# Set default timezone
ENV TZ=America/Los_Angeles

# Set default locale.
# /!\ locale-gen must be run as root.
RUN set -euo pipefail \
  && locale-gen en_US.UTF-8
ENV LANG=en_US.UTF-8
ENV LANGUAGE=en_US:en
ENV LC_ALL=en_US.UTF-8

# Install entrypoint script
COPY ./yolo_entrypoint.sh /
RUN set -euo pipefail \
  && chmod +x /yolo_entrypoint.sh

# Only for documentation purpose.
# Entrypoint and CMD are always set by the 
# Yolo agent when running the container.
ENTRYPOINT ["/yolo_entrypoint.sh"]
CMD ["sleep", "infinity"]

# Install the Docker CLI. 
# The Docker daemon socket will be mounted from instance.
RUN set -euo pipefail \
  && curl --fail --silent --show-error --location https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor --output /usr/share/keyrings/docker-archive-keyring.gpg \
  && echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release --codename --short) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null \
  && apt-get --assume-yes --quiet --quiet update \
  && apt-get --assume-yes --quiet --quiet install docker-ce-cli \
  && apt-get clean && rm --recursive --force /var/lib/apt/lists/* /tmp/*

# Install Docker compose
RUN set -euo pipefail \
  && LATEST_COMPOSE_VERSION=$(curl --fail --silent --show-error --location "https://api.github.com/repos/docker/compose/releases/latest" | grep --only-matching --perl-regexp '(?<="tag_name": ").+(?=")') \
  && curl --fail --silent --show-error --location "https://github.com/docker/compose/releases/download/${LATEST_COMPOSE_VERSION}/docker-compose-$(uname --kernel-name)-$(uname --machine)" --output /usr/libexec/docker/cli-plugins/docker-compose \
  && chmod +x /usr/libexec/docker/cli-plugins/docker-compose

# Install PHP
RUN set -euo pipefail \
  && apt-get --assume-yes --quiet --quiet update \
  && apt-get --assume-yes --quiet --quiet install \
    composer \
    php \
    php-all-dev \
    php-ctype \
    php-curl \
    php-date \
    php-gd \
    php-intl \
    php-json \
    php-mbstring \
    php-mysql \
    php-net-ftp \
    php-pgsql \
    php-php-gettext \
    php-sqlite3 \
    php-tokenizer \
    php-xml \
    php-zip \
  && apt-get clean && rm --recursive --force /var/lib/apt/lists/* /tmp/*

# Install Clang compiler (C/C++)
RUN set -euo pipefail \
  && curl --silent --show-error --location --fail https://apt.llvm.org/llvm-snapshot.gpg.key | apt-key add - \
  && apt-add-repository --yes "deb http://apt.llvm.org/jammy/ llvm-toolchain-jammy main" \
  && apt-get install --assume-yes --quiet --quiet \
    clang-format \
    clang-tools \
    cmake \
    clangd-14 \
  && update-alternatives --install /usr/bin/clangd clangd /usr/bin/clangd-14 100 \
  && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/*

# Install Java & Maven
RUN set -euo pipefail \
  && add-apt-repository --yes ppa:linuxuprising/java \
  && apt-get --assume-yes --quiet --quiet update \
  && echo oracle-java17-installer shared/accepted-oracle-license-v1-3 select true | debconf-set-selections \
  && apt-get install --assume-yes --quiet --quiet \
      gradle \
      oracle-java17-installer \
  && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/*

ARG MAVEN_VERSION=3.8.6
ENV MAVEN_HOME=/usr/share/maven
ENV PATH=$MAVEN_HOME/bin:$PATH
RUN set -euo pipefail \
  && mkdir --parents $MAVEN_HOME \
  && curl --silent --show-error --location --fail https://apache.osuosl.org/maven/maven-3/$MAVEN_VERSION/binaries/apache-maven-$MAVEN_VERSION-bin.tar.gz \
    | tar --extract --gzip --verbose --directory $MAVEN_HOME --strip-components=1

# Configure the user "yolo" in container.
# Triggered during build on instance.
# 
# We want the user "yolo" inside the container to get 
# the same permissions than the user "yolo" in the instance 
# (to access the Docker daemon, SSH keys and so on).
# 
# To do this, the two users need to share the same UID/GID.
RUN set -euo pipefail \
  && YOLO_USER_HOME_DIR="/home/yolo" \
  && YOLO_USER_WORKSPACE_DIR="${YOLO_USER_HOME_DIR}/workspace" \
  && YOLO_USER_WORKSPACE_CONFIG_DIR="${YOLO_USER_HOME_DIR}/.workspace-config" \
  && groupadd --gid 10000 --non-unique yolo \
  && useradd --gid 10000 --uid 10000 --non-unique --home "${YOLO_USER_HOME_DIR}" --create-home --shell /bin/bash yolo \
  && cp /etc/sudoers /etc/sudoers.orig \
  && echo "yolo ALL=(ALL) NOPASSWD:ALL" | tee /etc/sudoers.d/yolo > /dev/null \
  && groupadd --gid 10001 --non-unique docker \
  && usermod --append --groups docker yolo \
  && mkdir --parents "${YOLO_USER_WORKSPACE_CONFIG_DIR}" \
  && mkdir --parents "${YOLO_USER_WORKSPACE_DIR}" \
  && mkdir --parents "${YOLO_USER_HOME_DIR}/.ssh" \
  && mkdir --parents "${YOLO_USER_HOME_DIR}/.gnupg" \
  && mkdir --parents "${YOLO_USER_HOME_DIR}/.vscode-server" \
  && chown --recursive yolo:yolo "${YOLO_USER_HOME_DIR}" \
  && chmod 700 "${YOLO_USER_HOME_DIR}/.gnupg"

ENV USER=yolo
ENV HOME=/home/yolo
ENV EDITOR=/usr/bin/nano

ENV YOLO_WORKSPACE=/home/yolo/workspace
ENV YOLO_WORKSPACE_CONFIG=/home/yolo/.workspace-config

USER yolo
WORKDIR $HOME

# Install ZSH
RUN set -euo pipefail \
  && sudo apt-get --assume-yes --quiet --quiet update \
  && sudo apt-get --assume-yes --quiet --quiet install zsh \
  && sudo apt-get clean && sudo rm --recursive --force /var/lib/apt/lists/* /tmp/* \
  && mkdir .zfunc

# Install OhMyZSH and some plugins
RUN set -euo pipefail \
  && sh -c "$(curl --fail --silent --show-error --location https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" \
  && git clone --quiet https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions \
  && git clone --quiet https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting

# Change default shell for user "yolo"
RUN set -euo pipefail \
  && sudo usermod --shell $(which zsh) yolo

# Add a command "code" to ZSH.
# This command lets you open a file in VSCode 
# while being connected to an environment via SSH.
COPY --chown=yolo:yolo ./zsh/code_fn.zsh .zfunc/code

# Add .zshrc to home folder
COPY --chown=yolo:yolo ./zsh/.zshrc .

# Install Node.js.
# Nvm uses the "NODE_VERSION" env var to 
# choose which version needs to be installed.
ARG NODE_VERSION=18.7.0
ENV PATH=$PATH:$HOME/.nvm/versions/node/v$NODE_VERSION/bin
RUN set -euo pipefail \
  && curl --silent --show-error --location --fail https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash \
  && bash -c ". .nvm/nvm.sh \
    && npm config set python /usr/bin/python --global \
    && npm config set python /usr/bin/python \
    && npm install -g typescript \
    && npm install -g yarn"

# Install Python
ENV PYTHON_VERSION=3.10.6
ENV PATH=$PATH:$HOME/.pyenv/bin:$HOME/.pyenv/shims
RUN set -euo pipefail \
  && sudo apt-get --assume-yes --quiet --quiet update \
  && sudo apt-get --assume-yes --quiet --quiet install \
    libbz2-dev \
    libffi-dev \
    liblzma-dev \
    libncursesw5-dev \
    libreadline-dev \
    libsqlite3-dev \
    libxml2-dev \
    libxmlsec1-dev \
    llvm \
    make \
    tk-dev \
    xz-utils \
    zlib1g-dev \
  && sudo apt-get clean && sudo rm --recursive --force /var/lib/apt/lists/* /tmp/*
RUN set -euo pipefail \
  && curl --silent --show-error --location --fail https://github.com/pyenv/pyenv-installer/raw/master/bin/pyenv-installer | bash \
  && { echo; \
    echo 'eval "$(pyenv init -)"'; \
    echo 'eval "$(pyenv virtualenv-init -)"'; } >> .zshrc \
  && pyenv install ${PYTHON_VERSION} \
  && pyenv global ${PYTHON_VERSION} \
  && pip install virtualenv pipenv python-language-server[all]==0.19.0 \
  && rm -rf /tmp/*

# Install Ruby
ENV RUBY_VERSION=3.1.2
RUN set -euo pipefail \
  && curl --silent --show-error --location --fail https://rvm.io/mpapis.asc | gpg --import - \
  && curl --silent --show-error --location --fail https://rvm.io/pkuczynski.asc | gpg --import - \
  && curl --silent --show-error --location --fail https://get.rvm.io | bash -s stable \
  && bash -lc " \
    rvm requirements \
    && rvm install ${RUBY_VERSION} \
    && rvm use ${RUBY_VERSION} --default \
    && rvm rubygems current \
    && gem install bundler --no-document" \
  && { echo; \
    echo '[[ -s "$HOME/.rvm/scripts/rvm" ]] && . "$HOME/.rvm/scripts/rvm"'; } >> .zshrc

# Install Rust
RUN set -euo pipefail \
  && curl --proto '=https' --tlsv1.2 --silent --show-error --location --fail https://sh.rustup.rs | sh -s -- -y \
  && .cargo/bin/rustup update \
  && .cargo/bin/rustup component add rls-preview rust-analysis rust-src \
  && .cargo/bin/rustup completions zsh > ~/.zfunc/_rustup

# Install Go
ENV GO_VERSION=1.19
ENV PATH=$PATH:/usr/local/go/bin:$HOME/go/bin
RUN set -euo pipefail \
  && cd /tmp \
  && ARCH=$(arch | sed s/aarch64/arm64/ | sed s/x86_64/amd64/) \
  && curl --fail --silent --show-error --location "https://go.dev/dl/go${GO_VERSION}.linux-${ARCH}.tar.gz" --output go.tar.gz \
  && sudo tar --directory /usr/local --extract --file go.tar.gz \
  && rm go.tar.gz

WORKDIR $YOLO_WORKSPACE
```

As you can see, nothing fancy here. 

Yolo is built on `ubuntu 22.04` with most popular runtimes pre-installed. An user `yolo` is created and configured to be used as the default user. Root privileges are managed via `sudo`.

Your repositories will be cloned in `/home/yolo/workspace`. A default timezone and locale are set.

*(To learn more, see the [yolo-sh/base-env](https://github.com/yolo-sh/base-env) repository)*.

## Frequently asked questions

#### > Why using Docker as a VM and not something like NixOS, for example?

I'm aware that containers are not meant to be used as a VM like that but, at the time of writing, Docker is still the most widely used tool among developers to configure their environment (even if it may certainly change in the future).

#### > Given that my environment will run in a container does it mean that it will be limited?

Mostly not. 

Given the scope of this project (a private instance running in your own cloud provider account), Docker is mostly used for configuration purpose and not to "isolate" the VM from your environment.

As a result, your environment's container will run in **[privileged mode](https://docs.docker.com/engine/reference/run/#runtime-privilege-and-linux-capabilities)** in the **[host network](https://docs.docker.com/network/host/)**.

## License

Yolo is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).

<p align="center">
  <img src="https://user-images.githubusercontent.com/1233275/188123619-b37cd965-cdbf-4eb6-9218-fe81c136ed8a.jpeg" alt="yolo" width="180" height="180" />
</p>

<p align="center">
    <h1 align="center">Yolo</h1>
    <p align="center">Live environments for any repository. Running in your cloud provider account.<br />Currently available on <a href="https://github.com/yolo-sh/aws-cloud-provider">Amazon Web Services</a> and <a href="https://github.com/yolo-sh/hetzner-cloud-provider">Hetzner</a>.</p>
</p>

<blockquote align="left">
  ... use it to deploy your app, as a remote development environment or even as a code sandbox. Honestly, it's up to you!
</blockquote>

```bash
yolo aws init yolo-sh/api --instance-type t2.medium
```
<p align="center">
<img width="759" alt="Example of use of the Yolo CLI" src="https://user-images.githubusercontent.com/1233275/188150604-cdc1bcf7-301b-4756-9c87-153282a39532.png">
</p>
<blockquote align="left">
  ... now that this environment is created, you can connect to it with your preferred editor using the <code>edit</code> command:
</blockquote>

```bash
yolo aws edit yolo-sh/api
```

<img width="1136" alt="VSCode opened in a go repository" src="https://user-images.githubusercontent.com/1233275/188153382-99f30e14-b0db-444f-a956-c3e4cc44109f.png">


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
    - [Runtimes](#runtimes)
- [License](#license)

## Requirements

The Yolo binary has been tested on Linux and Mac. Support for Windows is theoretical ([testers needed](https://github.com/yolo-sh/cli/issues/1) 💙).

Before using Yolo, the following dependencies need to be installed:

- [OpenSSH Client](https://www.openssh.com/) (used to access your environments).

Before running the `edit` command, one of the following editors need to be installed:

- [Visual Studio Code](https://code.visualstudio.com/) (currently the sole editor supported).

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
yolo <cloud_provider> init <repository> [--instance-type=<instance_type>]
```
The `init` command initializes an environment for a specific GitHub repository.

An `--instance-type` flag can be passed to specify the instance type that will power your environment. (*See the corresponding cloud provider repository for default / valid values*).

#### Examples

```bash
yolo aws init yolo-sh/api
```

```bash
yolo hetzner init yolo-sh/api --instance-type cx11
```

### Edit

```bash
yolo <cloud_provider> edit <repository>
```
The `edit` command connects your preferred editor to an environment.

#### Example

```bash
yolo aws edit yolo-sh/api
```

### Open port

```bash
yolo <cloud_provider> open-port <repository> <port>
```
The `open-port` command makes the specified port reachable from any IP address using the `TCP` protocol.

#### Example

```bash
yolo aws open-port yolo-sh/api 8000
```

### Close port

```bash
yolo <cloud_provider> close-port <repository> <port>
```
The `close-port` command makes the specified port unreachable from all IP addresses.

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

All environments run in Docker containers built from the `ghcr.io/yolo-sh/workspace-full` image. You could see the source of this image in the [yolo-sh/workspace-full](https://github.com/yolo-sh/workspace-full) repository.

In summary, Yolo is built on `ubuntu 22.04` with `systemd` and most popular runtimes pre-installed. An user `yolo` is created and configured to be used as the default user. Root privileges are managed via `sudo`.

Your repositories will be cloned in `/home/yolo/workspace`.

Yolo uses the the [sysbox runtime](https://github.com/nestybox/sysbox) to improve isolation and to enable containers to run the same workloads than VMs.

### Runtimes

The following runtimes are pre-installed in all environments:

- `docker (latest)`

- `docker compose (latest)`

- `php (latest)`

- `java 17.0 / maven 3.8`

- `node 18.7 (via nvm)`

- `python 3.10 (via pyenv)`

- `ruby 3.1 (via rvm)`

- `rust (latest)`

- `go (latest)`

## License

Yolo is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).

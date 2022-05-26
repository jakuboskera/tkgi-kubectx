# `kubectx` with support of tkgi login

This repository is fork of [ahmetb/kubectx](https://github.com/ahmetb/kubectx).

There is added funcionallity of tkgi login for TKGI clusters.
So before switching to target context, tkgi login is done first.

## 📖 TOC
- [`kubectx` with support of tkgi login](#kubectx-with-support-of-tkgi-login)
  - [📖 TOC](#-toc)
  - [🤔 Purpose of `tkgi-kubectx` ?](#-purpose-of-tkgi-kubectx-)
  - [� Installation](#-installation)
    - [Homebrew (macOS and Linux)](#homebrew-macos-and-linux)
    - [Scoop (Windows)](#scoop-windows)
    - [From release](#from-release)
  - [🚀 Usage](#-usage)
    - [⚠️ Prerequisites](#️-prerequisites)
      - [Check if Kubernetes context exists in kubeconfig](#check-if-kubernetes-context-exists-in-kubeconfig)
      - [Installed tkgi CLI binary](#installed-tkgi-cli-binary)
    - [Create configuration files for `tkgi-kubectx`](#create-configuration-files-for-tkgi-kubectx)
      - [`~/.kube/tkgi-kubectx/config.yaml`](#kubetkgi-kubectxconfigyaml)
      - [`~/.kube/tkgi-kubectx/credentials.yaml`](#kubetkgi-kubectxcredentialsyaml)
    - [Test functionality](#test-functionality)
  - [🎉 Example usage](#-example-usage)
  - [ℹ️ Tips](#ℹ️-tips)
    - [Aliases](#aliases)
      - [macOS or Linux](#macos-or-linux)
      - [Windows](#windows)

## 🤔 Purpose of `tkgi-kubectx` ?

Original **kubectx** can just switch between contexts (clusters). But if user
works with TKGI clusters, he needs to first login to destination cluster before
switching a context.

Because of that functionality of tkgi login was added to **kubectx** and is called **tkgi-kubectx**.

It works almost the same as original **kubectx**. So you can utilize all the features
what original **kubectx** has, see more in original repository
[ahmetb/kubectx](https://github.com/ahmetb/kubectx).

## 🛠 Installation

### Homebrew (macOS and Linux)

If you use [Homebrew](https://brew.sh/) you can install like this:

```sh
brew install jakuboskera/tap/tkgi-kubectx
```

### Scoop (Windows)

If you use [Scoop](https://scoop.sh) you can install like this

```pwsh
scoop install https://raw.githubusercontent.com/jakuboskera/scoop-bucket/main/tkgi-kubectx.json
```

### From release

If you don't use any package managers mentioned above, you can install `tkgi-kubectx` binary manually.

1. Go to [Releases](https://github.com/jakuboskera/tkgi-kubectx/releases)
1. Download tar or zip file for your OS and CPU arch
1. Untar or unzip the file
1. Move binary `tkgi-kubectx` to your `PATH` (Check this manual [How to add a binary to your PATH](https://zwbetz.com/how-to-add-a-binary-to-your-path-on-macos-linux-windows))

## 🚀 Usage

### ⚠️ Prerequisites

1. Created a Kubernetes context in your kubeconfig for a which tkgi login will be performed
1. Installed `tkgi` CLI binary

#### Check if Kubernetes context exists in kubeconfig

To utilize kubectx with tkgi login feature you need to have already
created Kubernetes context in your kubeconfig for given cluster.

You can check if you have a context already created by

```bash
kubectl config get-contexts
```

In case previous command did not list a context for given cluster,
you have to do a manual login to the cluster first to create the
context in your kubeconfig.

In case previous command did list a context for given cluster, you are ok
and you can skip to next [section](#installed-tkgi-cli-binary).

There are two types of manual logins to TKGI clusters:

1. **cluster admin login** - if you are cluster admin

    ```bash
    tkgi login -a <TKGI API> -u <username> -k # -k if cert is self-signed
    ```

    ```bash
    tkgi get-credentials <cluster>
    ```

1. **non cluster admin login**  if you are not cluster admin

    ```bash
    tkgi get-kubeconfig <cluster> -a <TKGI API> -u <username> -k # -k if cert is self-signed
    ```

#### Installed tkgi CLI binary

Because `tkgi-kubectx` uses tkgi CLI binary under the hood, you have to have
tkgi CLI binary already installed.

You can check it by:

```bash
tkgi --version
```

If you don't have installed tkgi CLI binary, you can install from [HERE](https://docs.pivotal.io/tkgi/1-13/installing-cli.html).

### Create configuration files for `tkgi-kubectx`

If you met the listed prerequisities above, you can now configure `tkgi-kubectx`.

`tkgi-kubectx` needs for its function two files, these files must be created in specified location:

1. `~/.kube/tkgi-kubectx/config.yaml`
1. `~/.kube/tkgi-kubectx/credentials.yaml`

#### `~/.kube/tkgi-kubectx/config.yaml`

This file contains Kubernetes contexts, TKGI API and credentials reference.

```yaml
# ~/.kube/tkgi-kubectx/config.yaml

# contains list of clusters
clusters:
  # name of the cluster for a which tkgi login will be performed, the name is
  # usually same as name of the context
- name: <cluster1>
  # reference to username from ~/.kube/tkgi-kubectx/credentials.yaml
  creds: <username>
  # TKGI API URL for given cluster
  tkgiApi: https://<TKGI API>
- name: <clusterN>
  creds: <username>
  tkgiApi: https://<TKGI API>
...
```

#### `~/.kube/tkgi-kubectx/credentials.yaml`

This file contains list of credentials which are then used for login.

⚠️ Password is stored as plaintext, for this reason configuration was
split into two files.

```yaml
# ~/.kube/tkgi-kubectx/credentials.yaml

# contains list of credentials
credentials:
  # username is used as creds specified in ~/.kube/tkgi-kubectx/config.yaml
- username: <username>
  # plaintext password for <username>
  password: <password>
  # either true or false, because login commands are different
  # if user is cluster admin or not
  clusterAdmin: true
  # add how many users you want, the schema is same as for first username
- username: <another username>
  password: <password>
  clusterAdmin: false
...
```

### Test functionality

Now you should have everything configured and tkgi login will be performed
everytime before switching to context if this context is in `config.yaml`.

See example usage below.

## 🎉 Example usage

Let's say we have these three clusters:

| cluster (context) | TKGI API                      | username      | cluster admin |
| :---------------- | :---------------------------: | :-----------: | :-----------: |
| prod-cluster      | https://prod-tkgi.example.com | lhofstadter   | true          |
| test-cluster      | https://test-tkgi.example.com | rkoothrappali | false         |
| dev-cluster       | -                             | -             | -             |

For prod-cluster and test-cluster we need to perform tkgi login everytime
before switching to one of that contexts. For dev-cluster we don't need
a tkgi login as this cluster is for example local one.
For prod-cluster we will use user `lhofstadter` which is cluster admin
and for test-cluster we will use user `rkoothrappali` which is not cluster admin.

Check if we have created all contexts in our kubeconfig.

```bash
$ kubectl config get-contexts
dev-cluster
test-cluster
```

As you can see context for prod-cluster is missing, so we need to do
a manual login to create context for this cluster in our kubeconfig first.

Check if we have installed tkgi CLI binary

```bash
$ tkgi --version

TKGI CLI version: 1.13.0-build.212
```

As `lhofstadter` is cluster admin in prod-cluster,
we will use login commands for cluster admin:

```bash
tkgi login -a https://prod-tkgi.example.com -u lhofstadter -k # -k if cert is self-signed
```

```bash
tkgi get-credentials prod-cluster
```

Now check again if we have created all contexts in our kubeconfig.

```bash
$ kubectl config get-contexts
dev-cluster
prod-cluster
test-cluster
```

Well done. Now we will add credentials to `credentials.yaml`:

```yaml
# ~/.kube/tkgi-kubectx/credentials.yaml

credentials:
- username: lhofstadter
  password: ILoveSheldon
  clusterAdmin: true
- username: rkoothrappali
  password: iDontSpeakToGirls
  clusterAdmin: false
```

Now we will add configuration for prod-cluster and test-cluster
to `config.yaml`:

```yaml
# ~/.kube/tkgi-kubectx/config.yaml

clusters:
- name: prod-cluster
  creds: lhofstadter
  tkgiApi: https://prod-tkgi.example.com
- name: test-cluster
  creds: rkoothrappali
  tkgiApi: https://test-tkgi.example.com
```

Everything should be configured now. Let's try it.
If we switch to prod-cluster, it will use cluster login commands and user
`lhofstadter` to login to that cluster.

```bash
$ tkgi-kubectx prod-cluster
Login first before switching context to "prod-cluster"...

API Endpoint: https://prod-tkgi.example.com
User: lhofstadter
Login successful.


Fetching credentials for cluster prod-cluster.
PKS_USER_PASSWORD is set. Reading password from PKS_USER_PASSWORD
Context set for cluster prod-cluster.

You can now switch between clusters by using:
$kubectl config use-context <cluster-name>

✔ Switched to context "prod-cluster".
```

Now when we switch to test-cluster, it will use user `rkoothrappali`.
As this user is not cluster admin, it will use non cluster admin login command.

```bash
$ tkgi-kubectx test-cluster
Login first before switching context to "test-cluster"...

Fetching kubeconfig for cluster test-cluster and user rkoothrappali.
You can now use the kubeconfig for user rkoothrappali:
$kubectl config use-context test-cluster

✔ Switched to context "test-cluster".
```

Cluster dev-cluster is not specified in `config.yaml` so for this context,
it will not perform tkgi login, it will just switch to that context.

```bash
$ tkgi-kubectx dev-cluster
✔ Switched to context "dev-cluster".
```

## ℹ️ Tips

### Aliases

Name `tkgi-kubectx` is relatively long for a command so you can create an alias for it

#### macOS or Linux

1. bash

    ```bash
    echo "alias kx=tkgi-kubectx" >> ~/.bashrc
    ```

1. zsh

    ```bash
    echo "alias kx=tkgi-kubectx" >> ~/.zshrc
    ```

#### Windows

1. PowerShell

    ```pwsh
    Set-Alias -Name kx -Value tkgi-kubectx
    ```

# `kubectx` with support of `tkgi login`

This repository is fork of [ahmetb/kubectx](https://github.com/ahmetb/kubectx).

There is added funcionallity of `tkgi login` for TKGI clusters.
So before switching to target context, tkgi login is done first.

## Purpose of `tkgi-kubectx` ?

Original **kubectx** can just switch between contexts (clusters). But if user
works with TKGI clusters, he needs to first login to destination cluster before
switching a context.

From that reason, funcionallity of `tkgi login` was added to **kubectx** and was it
named as **tkgi-kubectx**.

It works almost the same as original **kubectx**. So you can utilize all the features
what original **kubectx** has, see more in original repository
[ahmetb/kubectx](https://github.com/ahmetb/kubectx).

## How to use a `tkgi-kubectx`

Usage `tkgi-kubectx` is really almost the same as original `kubectx`.
Hovewer if you create files `~/.kube/tkgi-kubectx/config.yaml` and `~/.kube/tkgi-kubectx/credentials.yaml`,
`tkgi-kubectx` will start using login feature automatically.

### Prerequisities

1. installed [tkgi](https://docs.pivotal.io/tkgi/1-13/installing-cli.html)
1. created a file `~/.kube/tkgi-kubectx/config.yaml` - contains clusters name, TGKI API URL for that cluster and username
1. created a file `~/.kube/tkgi-kubectx/credentials.yaml` - contains username, password and flag if user is cluster admin or not

#### Create config and credentials files

1. Create a folder `tkgi-kubectx` in `~/.kube`

    ```bash
    mkdir -p ~/.kube/tkgi-kubectx
    ```

1. Create a file `~/.kube/tkgi-kubectx/config.yaml`

    ```bash
    touch ~/.kube/tkgi-kubectx/config.yaml
    ```

1. Add content to the file

    ```yaml
    # ~/.kube/tkgi-kubectx/config.yaml

    clusters:
      # name of the cluster for a which tkgi login will be performed, the name is
      # usually same as name of the context
    - name: <cluster1>
      # this username must be in ~/.kube/tkgi-kubectx/credentials.yaml
      creds: <username>
      # TKGI API URL where you normally login
      tkgiApi: https://<tkgi api url>
      # add how many clusters you want, the schema is same as for first cluster
    - name: <clusterN>
      creds: <username>
      tkgiApi: https://<tkgi api url>
    ...
    ```

1. Create a file `~/.kube/tkgi-kubectx/credentials.yaml`

    ```bash
    touch ~/.kube/tkgi-kubectx/credentials.yaml
    ```

1. Add content to the file

    ```yaml
    # ~/.kube/tkgi-kubectx/credentials.yaml

    credentials:
      # username specified in ~/.kube/tkgi-kubectx/config.yaml
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

1. From now on everytime you switch context to different context and that context is configured
   in  `~/.kube/tkgi-kubectx/config.yaml`, it will first login to target context (cluster).
   Also when target context is using same TKGI API URI as previous contecxt, login is skipped,
   as it is not needed and it is faster way to switch.

## Installation

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
1. Move binary `tkgi-kubectx` to your `PATH`

## Tips

Name `tkgi-kubectx` is relatively long name for a command so you can create an alias for it

### MacOS or Linux

```bash
echo "alias kx=tkgi-kubectx" >> ~/.bashrc # or ~/.zshrc if you use ZSH
```

### Windows

```pwsh
Set-Alias -Name kx -Value tkgi-kubectx
```
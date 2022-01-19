# Installation

> Guides for `Arch`, `Fedora`, `brew` and many more coming soon!!

## Download binaries

Go to [the releases page](https://github.com/bare-cli/bare/releases), find the version you want, and download the zip file. Unpack the zip file, and put the binary to somewhere you want (on UNIX-y systems, /usr/local/bin or the like). Make sure it has execution bits turned on.

## For Ubuntu/Debian based systems

```shell
wget https://github.com/bare-cli/bare/releases/download/v0.1.8/bare_0.1.8_linux_amd64.deb
sudo dpkg -i bare_0.1.8_linux_amd64.deb
```

## Build from source

- Clone the repo

```shell
git clone https://github.com/bare-cli/bare.git
```

using [gh](https://github.com/cli/cli)

```shell
gh repo clone bare-cli/bare <destination>
```

- Change directory to your local copy:

```shell
cd <folder-name>
```

- Install all the go dependencies

```shell
make setup
```

- Build `Bare`

```shell
make build
```

The output binary can be found in `bin` folder

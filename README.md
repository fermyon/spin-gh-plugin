# Spin GitHub Plugin

This is a plugin that generates GitHub Actions for your Spin Apps.

# Installation

## Install the latest version of the plugin

The latest stable release of the `gh` plugin can be installed like so:

```sh
spin plugins update
spin plugin install gh
```

## Install the canary version of the plugin

The `canary` release of the `gh` represents the most recent commits on `main` and may not be stable, with some features still in progress.

```sh
spin plugins install --url https://github.com/ThorstenHans/spin-gh-plugin/releases/download/canary/gh.json
```

## Install from a local build

Alternatively, use the `spin pluginify` plugin to install from a fresh build. This will use the pluginify manifest (`spin-pluginify.toml`) to package the plugin and proceed to install it:

```sh
spin plugins install pluginify
go build -o gh main.go
spin pluginify --install
```

# Usage

## The `create-action` command

To create a new GitHub Action for your Spin App(s), execute the following command:

```bash
spin gh create-action
```

### Arguments & Flags

The `create-action` command accepts a bunch of commands that you can use to customize its behavior:

#### GitHub Action Triggers

| Argument | Alias | Description | Default |
|----------|-------|-------------|---------|
| `ci` | | Run GitHub Action for every push on the specified branch (CI) | `main` |
| `cron` | | Run GitHub Action on a Cron Schedule | |
| `pr` | | Run GitHub Action for every PR targeting the specified branch | |
| `manual` | | Run GitHub Action on manual dispatch | `false` |


#### Tool Versions

You can use the following arguments to customize versions installed as part of the GitHub Action

| Argument | Alias | Description | Default |
|----------|-------|-------------|---------|
| `spin-version` | | Pin the Spin Version | latest stable release |
| `rust-version` | | Pin the Rust Version | `1.80.1` |
| `go-version` | | Pin the Go Version | `1.23.2` |
| `tinygo-version` | | Pin the TinyGo Version | `0.33.0` |
| `node-version` | | Pin the Node.js Version | `22` |
| `python-version` | | Pin the Python Version | `3.13.0` |

#### General GitHub Action customization

| Argument | Alias | Description | Default |
|----------|-------|-------------|---------|
| `env` | | Specify Environment Variables (format key=value) | |
| `name` | `n` | Specify the name of the GitHub Action | `CI` |
| `plugin` | `p` | Specify Spin Plugins that should be installed | |
| `os` | | Specify the operating system | `ubuntu-latest` |

### Render Options

| Argument | Alias | Description | Default |
|----------|-------|-------------|---------|
| `output` | `o` | Output path | `./github/workflows/ci.yaml` |
| `dry-run` | | Print GitHub Action yaml to `stdout` | `false` |
| `overwrite` | | Overwrite the output if it exists | `false` |
| `template` | `t` | Provide a custom template | |

## The `eject` command

TBD

### Template Variables

TBD
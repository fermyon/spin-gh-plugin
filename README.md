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

TBD
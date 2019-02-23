# Rapid Compose (RC)

A teminal UI to start individual services from docker-copose.yaml

The `rc` looks for `docker-compose.yaml` or `docker-compose.yml`

## Install

Homebrew:

```sh
brew install trusz/tap/rapid-compose
```

## Usage

Go into a folder where there is a `docker-compose` file and execute `rc`

By default, `rc` filters out every dependency services (`depends_on`) in order to show only the main ones.  

### Options

- `-a`: Show all services
- `-i`: Inverse selection. Start everything except selected ones.
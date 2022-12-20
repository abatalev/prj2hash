# Prj2Hash -- calculate hash for project

## Usage

```sh
$ cd project
$ prj2hash
08f5a690f079a1fae4596ee61f4ab9fed24666f0
```

## Config 

```yaml
rules:
- deny **/*
- allow **/*.go
```

DEPRECATED! Support of section `excludes` will be removed in the next version

```yaml
excludes:
- target/**/*
- readme.md
```

Rewrite section `excludes` using `rules`.

```yaml
rules:
- allow **/*
- deny target/**/*
- deny readme.md
```

## Build

### Simple build
```sh
go mod tidy
go build .
```

### Build with vendor and docker

```sh
go mod vendor
docker build -t abatalev/prj2hash .
```

## Development

```sh
./build.sh
```
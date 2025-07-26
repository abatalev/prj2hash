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

## Examples

```sh
❯ ./prj2hash --dry-run --cfg examples/nginxs/nginx1-prj2hash.yaml examples/nginxs
 file Dockerfile.nginx1 cd19de9379c5e37ae84eb1853e5177cb7cf3420e
 file nginx1/index.html 9b895ef2c4a85a7ea565ef768504115e5474ac00
 file script.js 474660db527afb1bd7e5f5418b24bd63ea6f34ed
total a882780f a882780f89bdd4e28ca222daa963ae56dd3cf675
❯ ./prj2hash --dry-run --cfg examples/nginxs/nginx2-prj2hash.yaml examples/nginxs
 file Dockerfile.nginx2 8d23f104e2631b3b2af28c891116c36cf636aa58
 file nginx2/index.html f601fa1fc00f12c511ece27e82fba829e4e64251
 file script.js 474660db527afb1bd7e5f5418b24bd63ea6f34ed
total 139d40be 139d40be110804c7bc97f3372def074d419e21c2
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
docker buildx build -f Dockerfile.alpine -t prj2hash:latest .
```

or

```sh
./build.sh
```
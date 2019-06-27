# Redirector

Takes any request and redirects it to `$TARGET`

uses 302 StatusFound by default

## Usage

```sh
env TARGET=https://www.someplace.com bin/redirector
```

Permanent Redirect

```sh
env TARGET=https://www.someplace.com REDIRECT=301 bin/redirector
```

## Dev

```sh
ag -g '\.go' . | entr sh -c 'clear && env PORT=8080 TARGET=https://www.someplace.com make dev'
```

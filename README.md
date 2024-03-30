# Yokai Worker Template

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go version](https://img.shields.io/badge/Go-1.22-blue)](https://go.dev/)

> Worker application template based on the [Yokai](https://github.com/ankorstore/yokai) Go framework.

<!-- TOC -->
* [Documentation](#documentation)
* [Overview](#overview)
  * [Layout](#layout)
  * [Makefile](#makefile)
* [Getting started](#getting-started)
  * [Installation](#installation)
    * [With GitHub](#with-github)
    * [With gonew](#with-gonew)
  * [Usage](#usage)
<!-- TOC -->

## Documentation

For more information about the [Yokai](https://github.com/ankorstore/yokai) framework, you can check its [documentation](https://ankorstore.github.io/yokai).

## Overview

This template provides:

- a ready to extend [Yokai](https://github.com/ankorstore/yokai) application, with the [worker](https://ankorstore.github.io/yokai/modules/fxworker/) module installed
- a ready to use [dev environment](docker-compose.yaml), based on [Air](https://github.com/cosmtrek/air) (for live reloading)
- a ready to use [Dockerfile](Dockerfile) for production
- some examples of [worker](internal/worker/example.go) and [test](internal/worker/example_test.go) to get started

### Layout

This template is following the [recommended project layout](https://go.dev/doc/modules/layout#server-project):

- `cmd/`: entry points
- `configs/`: configuration files
- `internal/`:
  - `worker/`: worker and test examples
  - `bootstrap.go`: bootstrap
  - `register.go`: dependencies registration

### Makefile

This template provides a [Makefile](Makefile):

```
make up     # start the docker compose stack
make down   # stop the docker compose stack
make logs   # stream the docker compose stack logs
make fresh  # refresh the docker compose stack
make test   # run tests
make lint   # run linter
```

## Getting started

### Installation

#### With GitHub

You can create your repository [using the GitHub template](https://github.com/new?template_name=yokai-worker-template&template_owner=ankorstore).

It will automatically rename your project resources and push them, this operation can take a few minutes.

Once ready, after cloning and going into your repository, simply run:

```shell
make fresh
```

#### With gonew

You can install [gonew](https://go.dev/blog/gonew), and simply run:

```shell
gonew github.com/ankorstore/yokai-worker-template github.com/foo/bar
cd bar
make fresh
```

### Usage

Once ready, the application core dashboard will be available on [http://localhost:8081](http://localhost:8081).

To see the [provided example worker](internal/worker/example.go) in action, simply run:

```shell
make logs
```

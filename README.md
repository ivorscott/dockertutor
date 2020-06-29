# dockertutor

## Proof Of Concept (WIP)

### Introduction

Beginners need tutorials for docker. To get one from docker currently they need to go to the official documentation. What if these tutorials or something similar were added directly to the cli.

[GITHUB ISSUE](https://github.com/docker/roadmap/issues/102)

### Usage

dockertutor has 3 tutorials:

1. docker
2. docker-compose
3. swarm

First initialize your workspace then switch between tutorials with `--category` or `-c`

```
go build
go install

dockertutor init
dockertutor -c docker
```

### Tests

Run tests

```
go tests -v ./...
```

### UX Issues

Currently digging into ways to improve the user experience when working with lessons that require example files and file modification, in addition to running docker commands against them. My latest merge is getting close. https://github.com/ivorscott/dockertutor/pull/16

### Contact

ivor@devpie.io

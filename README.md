# dockertutor

## Proof Of Concept (WIP)

### Introduction

Beginners need tutorials for docker. To get one from docker currently they need to go to the official documentation. What if these tutorials or something similar were added directly to the cli.

[GITHUB ISSUE](https://github.com/docker/roadmap/issues/102)

### Go Docs
Run the following command in the project root.
```
docker run -v $(pwd):/go/src -p 8080:8080 ivorsco77/godocs  
```
Then navigate to http://localhost:8080/pkg/tutor/

### Usage
dockertutor has 3 tutorials:

1) docker
2) docker-compose
3) swarm

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

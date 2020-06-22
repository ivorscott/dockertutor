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

### Demo
dockertutor has 3 tutorials:

1) docker
2) docker compose
3) swarm

-c stands for category. The first command defaults to docker cli
```
./dockertutor  # defaults to the "docker" category
./dockertutor -c docker-compose # WIP, requires file generation and modification
./dockertutor -c swarm 
```

### Tests 

Run tests
```
go tests -v ./...
```
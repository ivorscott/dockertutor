# dockertutor

## Proof Of Concept (WIP)

### Introduction

Beginners need tutorials for docker. To get one from docker currently they need to go to the official documentation. What if these tutorials or something similar were added directly to the cli.

[GITHUB ISSUE](https://github.com/docker/roadmap/issues/102)

### Go Docs
Clone the repository. In the project directory run the following command and then navigate to
http://localhost:8080/pkg/#thirdparty. Click on tutor.
```
docker run -v $(pwd):/go/src -p 8080:8080 ivorsco77/godocs  
```

### Demo

```
go build ./cmd/dockertutor
./dockertutor
```

### Tests 

Run tests
```
go tests -v ./...
```
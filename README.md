# dockertutor

## Proof Of Concept (WIP)

### Introduction

Beginners need tutorials for docker. To get one from docker currently they need to go to the official documentation. What if these tutorials or something similar were added directly to the cli.

[GITHUB ISSUE](https://github.com/docker/roadmap/issues/102)

### Demo

```
go build ./cmd/dockertutor
./dockertutor
```

### TODO

[ ] User is required to enter a project working directory (for generated files)

[ ] User can choose between 3 tutorials (docker cli, docker-compose, swarm)

[ ] Each tutorial contains a number of lessons

[ ] Each lesson renders an instruction

[ ] Each lesson prompts a call to action

[ ] Docker commands are required to move to the next lesson

[ ] Some lessons generate files in the user's selected working directory

[ ] Some lessons require the user to modify files and a docker command to move to the next lesson

[ ] Some lessons reveal explanations before moving to the next lesson

[ ] Successfully completing a lesson is rewarded

[ ] Multiple tutorial runs with the same working directory doesn't overwrite previous files

[ ] User can quit the program and return to the same state

[ ] User receives a progress bar indicating tutorial completion

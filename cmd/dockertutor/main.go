package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/ivorscott/dockertutor/internal/lessons"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	fmt.Print(lessons.DockerIntro)
	cmdStr := "docker run hello-world"
	out, err := exec.Command("/bin/sh", "-c", cmdStr).CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Printf("%s", out)

	return nil
}

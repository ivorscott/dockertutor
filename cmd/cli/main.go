package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/ivorscott/dockertutor/internal/text"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	fmt.Print(text.Introduction)
	fmt.Print(text.Linebreak)
	fmt.Print(text.Lesson_1_1_Text)
	cmdStr := "docker run hello-world"
	out, err := exec.Command("/bin/sh", "-c", cmdStr).CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Printf("%s", out)

	return nil
}

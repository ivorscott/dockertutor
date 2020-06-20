package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/ivorscott/dockertutor/tutor"
)

var categories = [3]string{"docker", "docker-compose", "swarm"}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func prompt(stdin io.Reader) (string, error) {
	fmt.Print("> ")
	reader := bufio.NewReader(stdin)
	return reader.ReadString('\n')
}

func run() error {
	cat := flag.String("c", categories[0], "Select tutorial category")
	flag.Parse()

	t, err := tutor.NewTutorial(*cat)
	if err != nil {
		return err
	}

	if t.ActiveLessonId == 0 {
		t.Welcome()
	}

	t.ActiveLesson.Explain()

	cmd, err := prompt(os.Stdin)
	if err != nil {
		return err
	}

	out, err := exec.Command("/bin/sh", "-c", cmd).CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Printf("%s", out)
	return nil
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/ivorscott/dockertutor/tutor"
)

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
	cat := flag.String("c", tutor.Categories[0], "Select tutorial category")
	flag.Parse()

	tutsConf := "./tutor/tutorials.json"
	lessConf := fmt.Sprintf("./tutor/%s.json", *cat)

	tutsData, err := ioutil.ReadFile(tutsConf)
	if err != nil {
		return err
	}

	lessData, err := ioutil.ReadFile(lessConf)
	if err != nil {
		return err
	}

	t, err := tutor.NewTutorial(tutsData, lessData, *cat)
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

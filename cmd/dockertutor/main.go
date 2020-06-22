package main

import (
	"flag"
	"fmt"
	"github.com/ivorscott/dockertutor/tutor"
	"io/ioutil"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
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

	t.NextLesson()

	return nil
}

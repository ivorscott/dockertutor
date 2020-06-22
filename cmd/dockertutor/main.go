package main

import (
	"flag"
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

	tutsConfig, lessonConfig := tutor.ConfigFiles(*cat)

	tutsData, err := ioutil.ReadFile(tutsConfig)
	if err != nil {
		return err
	}

	lessData, err := ioutil.ReadFile(lessonConfig)
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

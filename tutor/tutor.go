package tutor

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// Tutorial manages active tutorial state
type Tutorial struct {
	Category       string
	ActiveLessonId int
	ActiveLesson   Lesson
	Lessons
}

type Tutorials []Tutorial

// Lesson represents the state of an exercise
type Lesson struct {
	Id          int
	Title       string
	Exercise    string
	Answer      string
	Examples    []string
	Explanation string
	Complete    bool
	Resources
}

// Lessons represent the Tutorial lessons
type Lessons []Lesson

// Resources aim to preserve the state of a lesson in the event of
// exiting the program and starting back up again. Direct resources
// and the resources of dependent lessons
type Resources struct {
	Images     []string
	Containers []string
	Volumes    []string
	Networks   []string
}

var IntroMap = map[string]string{
	"docker":         DockerIntro,
	"swarm":          SwarmIntro,
	"docker-compose": ComposeIntro,
}

var CatMap = map[string]int{
	"docker":         0,
	"swarm":          2,
	"docker-compose": 1,
}

const config = "tutor/tutorials.json"

var ErrFileEmpty = errors.New("config file is empty")
var ErrFileNotRead = errors.New("config file could not be read")

// NewTutorial returns a new tutorial by category
func NewTutorial(category string) (*Tutorial, error) {
	t := &Tutorials{}
	file, err := ioutil.ReadFile(config)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, os.ErrNotExist
		}
		return nil, ErrFileNotRead
	}

	if len(file) == 0 {
		return nil, ErrFileEmpty
	}

	if err := json.Unmarshal(file, t); err != nil {
		return nil, err
	}

	l, err := NewLessons(category)
	if err != nil {
		return nil, err
	}

	tuts := *t
	cat := CatMap[category]
	al := *l

	return &Tutorial{
		Category:       category,
		ActiveLessonId: tuts[cat].ActiveLessonId,
		ActiveLesson:   al[tuts[cat].ActiveLessonId],
		Lessons:        *l,
	}, nil
}

func NewLessons(category string) (*Lessons, error) {
	l := &Lessons{}
	lSrc := fmt.Sprintf("tutor/%s.json", category)
	file, err := ioutil.ReadFile(lSrc)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, os.ErrNotExist
		}
		return nil, ErrFileNotRead
	}

	if len(file) == 0 {
		return nil, ErrFileEmpty
	}

	if err := json.Unmarshal(file, l); err != nil {
		return nil, err
	}

	return l, nil
}

// Tutorial returns a lesson exercise
func (t *Tutorial) Welcome() {
	fmt.Fprintln(os.Stdout, IntroMap[t.Category])
}

// Next fetches the next lesson
func (t *Tutorial) Next() {}

// Setup provisions lesson resources
func (l *Lesson) Setup() {}

// Teardown tears down lesson resources
func (l *Lesson) Teardown() {}

// Teach returns a lesson exercise
func (l *Lesson) Teach() {}

// Explain returns a lesson explanation
func (l *Lesson) Explain() {
	fmt.Fprintln(os.Stdout, l.Exercise)
}

// Success represents a lesson succeeded
func (l *Lesson) Success() {}

// Failure represents a lesson failed
func (l *Lesson) Failure() {}

// Exit quits the lesson
func (l *Lesson) Exit() {}

// Reset resources and lesson progress
func (l *Lesson) Reset() {}

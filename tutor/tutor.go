// Tutor provides state management for tutorials and their lessons
package tutor

import (
	"encoding/json"
	"fmt"
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

var Categories = [3]string{"docker", "docker-compose", "swarm"}

var IntroMap = map[string]string{
	Categories[0]: dockerIntro(),
	Categories[1]: composeIntro(),
	Categories[2]: swarmIntro(),
}

var catMap = map[string]int{
	Categories[0]: 0,
	Categories[1]: 2,
	Categories[2]: 1,
}

// NewTutorial returns a new tutorial by category
func NewTutorial(tutsData, lessData []byte, category string) (*Tutorial, error) {
	t := &Tutorials{}

	if err := json.Unmarshal(tutsData, t); err != nil {
		return nil, err
	}

	l, err := NewLessons(lessData)
	if err != nil {
		return nil, err
	}

	tuts := *t
	cat := catMap[category]
	al := *l

	return &Tutorial{
		Category:       category,
		ActiveLessonId: tuts[cat].ActiveLessonId,
		ActiveLesson:   al[tuts[cat].ActiveLessonId],
		Lessons:        *l,
	}, nil
}

func NewLessons(lessData []byte) (*Lessons, error) {
	l := &Lessons{}

	if err := json.Unmarshal(lessData, l); err != nil {
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

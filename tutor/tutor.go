// Tutor provides state management for tutorials and their lessons
package tutor

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
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

func Prompt(stdin io.Reader) (string, error) {
	fmt.Print("\n> ")
	reader := bufio.NewReader(stdin)
	return reader.ReadString('\n')
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

// New Lessons provides the lessons state for the tutorial
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
func (t *Tutorial) NextLesson() {
	var answer = false

	for answer == false {
		t.ActiveLesson.Teach()

		cmd, err := Prompt(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err.Error())
		}

		answer = t.CheckAnswer(cmd)

		out, err := exec.Command("/bin/sh", "-c", cmd).CombinedOutput()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err.Error())
		}
		fmt.Printf("%s", out)

		if answer {
			t.Success()
			return
		} else {
			t.Failure()
		}
	}

}

// CheckAnswer splits the command on a newline and checks the answer
func (t *Tutorial) CheckAnswer(cmd string) bool {
	answer := bytes.Split([]byte(cmd), []byte("\n"))

	if bytes.Equal(answer[0], []byte(t.ActiveLesson.Answer)) {
		return true
	}
	return false
}

// Setup provisions lesson resources
func (l *Lesson) Setup() {}

// Teardown tears down lesson resources
func (l *Lesson) Teardown() {}

// Teach returns a lesson exercise
func (l *Lesson) Teach() {
	fmt.Fprintln(os.Stdout, l.Title)
	fmt.Fprintln(os.Stdout, lbreak())
	fmt.Fprintln(os.Stdout, l.Exercise)
}

// Explain returns a lesson explanation
func (l *Lesson) Explain() {
	fmt.Fprintln(os.Stdout, l.Explanation)
}

// Success represents a lesson succeeded
func (t *Tutorial) Success() {
	fmt.Fprintln(os.Stdout, "Correct!")
	fmt.Println()
	t.ActiveLesson = t.Lessons[t.ActiveLessonId+1]
	t.ActiveLessonId = t.ActiveLessonId + 1
	t.NextLesson()
}

// Failure represents a lesson failed
func (t *Tutorial) Failure() {
	fmt.Fprintln(os.Stdout, "Command was not correct.")
	fmt.Println()
}

// Exit quits the lesson
func (l *Lesson) Exit() {}

// Reset resources and lesson progress
func (l *Lesson) Reset() {}

package tutor

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

// Lesson represents the state of an exercise
type Lesson struct {
	Title       string
	Exercise    string
	Answer      string
	Example     string
	Explanation string
	Complete    bool
	Setup       []string
	Teardown    []string
}

type Lessons []Lesson

// New Lessons provides the lessons state for the tutorial
func NewLessons(lessData []byte) (*Lessons, error) {
	l := &Lessons{}
	if err := json.Unmarshal(lessData, l); err != nil {
		return nil, err
	}
	return l, nil
}

// Setup provisions lesson resources
func (l *Lesson) setup() error {
	if len(l.Setup) > 0 {
		for _, cmd := range l.Setup {
			c := exec.Command("/bin/sh", "-c", cmd)
			if c != nil {
				err := c.Start()
				if err != nil {
					return err
				}
				err = c.Wait()
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// Teardown tears down lesson resources
func (l *Lesson) teardown() error {
	if len(l.Teardown) > 0 {
		for _, cmd := range l.Teardown {
			c := exec.Command("/bin/sh", "-c", cmd)
			if c != nil {
				err := c.Start()
				if err != nil {
					return err
				}
				c.Wait()
				// ignoring teardown errors
			}
		}
	}
	return nil
}

// Teach returns a lesson exercise
func (l *Lesson) teach() error {
	if _, err := fmt.Fprintf(os.Stdout, "\n%s\n%s\n%s", l.Title, lbreak(), l.Exercise); err != nil {
		return err
	}
	return nil
}

// Explain returns a lesson explanation
func (l *Lesson) explain() error {
	if _, err := fmt.Fprintln(os.Stdout, l.Explanation); err != nil {
		return err
	}
	return nil
}

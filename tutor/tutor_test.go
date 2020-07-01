package tutor

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

const (
	dockerLess  = "./testdata/docker.json"
	composeLess = "./testdata/docker-compose.json"
	swarmLess   = "./testdata/swarm.json"
)

func TestPrompt(t *testing.T) {
	want := "docker run hello-world"
	reader := bufio.NewReader(strings.NewReader(want))

	got, err := prompt(reader)
	if err != io.EOF {
		t.Errorf("Prompt failed and sent: %s", err)
	}

	t.Run("input matches output", func(t *testing.T) {
		if !bytes.Equal([]byte(got), []byte(want)) {
			t.Fatalf("Expected %s, got %s instead", want, got)
		}
	})
}

func TestNewTutorial(t *testing.T) {
	tests := []struct {
		Category   string
		LessonFile string
	}{
		{"docker", dockerLess},
		{"docker-compose", composeLess},
		{"swarm", swarmLess},
	}

	config, err := os.Open("./testdata/tutor_config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer config.Close()

	c, err := ioutil.ReadAll(config)
	if err != nil {
		log.Fatal(err)
	}

	for _, tt := range tests {
		lf, err := os.Open(tt.LessonFile)
		if err != nil {
			t.Fatal(err)
		}

		l, err := ioutil.ReadAll(lf)
		if err != nil {
			t.Fatal(err)
		}

		tut, err := NewTutorial(c, l, tt.Category)
		if err != nil {
			t.Error(err)
		}

		t.Run(tt.Category, func(t *testing.T) {
			t.Run("should have category", func(t *testing.T) {
				if tut.Category != tt.Category {
					t.Errorf("Category not set correctly")
				}
			})
			t.Run("should have active lesson", func(t *testing.T) {
				if &tut.ActiveLesson == &tut.Lessons[tut.ActiveLessonIndex] {
					t.Errorf("Active Lesson is not set correctly")
				}
			})
		})

		lf.Close()
	}
}

func TestNewLessons(t *testing.T) {
	tests := []struct {
		Category   string
		LessonFile string
	}{
		{"docker", dockerLess},
		{"docker-compose", composeLess},
		{"swarm", swarmLess},
	}

	for _, tt := range tests {
		lessData, err := ioutil.ReadFile(tt.LessonFile)
		if err != nil {
			t.Fatal(err)
		}

		less, err := NewLessons(lessData)
		if err != nil {
			t.Error(err)
		}

		t.Run(tt.Category, func(t *testing.T) {
			for _, lt := range *less {
				t.Run("should have Title", func(t *testing.T) {
					if lt.Title == "" {
						t.Errorf("Title is empty")
					}
				})
				t.Run("should have Exercise", func(t *testing.T) {
					if lt.Exercise == "" {
						t.Errorf("Excerice is empty")
					}
				})
				t.Run("should have Answer", func(t *testing.T) {
					if lt.Answer == "" {
						t.Errorf("Answer is empty")
					}
				})
			}
		})
	}
}

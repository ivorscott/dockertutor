// Tutor provides state management for tutorials and their lessons
package tutor

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// Tutorial manages active tutorial state
type Tutorial struct {
	Category       string
	Directory      string
	ActiveLessonId int
	ActiveLesson   Lesson `json:"-"`
	Tutorials      `json:"-"`
	Lessons        `json:"-"`
}

type Tutorials []Tutorial

// Lesson represents the state of an exercise
type Lesson struct {
	Id          int
	Title       string
	Exercise    string
	Answer      string
	Example     string
	Explanation string
	Complete    bool
	Setup       []string
	Teardown    []string
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
	Categories[1]: 1,
	Categories[2]: 2,
}

func prompt(stdin io.Reader) (string, error) {
	fmt.Print("\n> ")
	reader := bufio.NewReader(stdin)
	return reader.ReadString('\n')
}

func ConfigFiles(category string) (string, string) {
	tutsConfig := "./tutor/tutorials.json"
	lessonConfig := fmt.Sprintf("./tutor/%s.json", category)
	return tutsConfig, lessonConfig
}

func templatePath() string {
	return `./examples`
}

// NewTutorial returns a new tutorial by category
func NewTutorial(tutsData, lessData []byte, category, directory string) (*Tutorial, error) {
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

	tut := &Tutorial{
		Category:       category,
		Directory:      tuts[cat].Directory,
		ActiveLessonId: tuts[cat].ActiveLessonId,
		ActiveLesson:   al[tuts[cat].ActiveLessonId],
		Tutorials:      tuts,
		Lessons:        *l,
	}
	if directory == "" {
		if tut.Directory == "" {
			flag.Usage()
			os.Exit(1)
		}
	} else {
		tut.Directory = directory
	}

	tut.save()
	return tut, nil
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
	if cmd := t.ActiveLesson.teardown(); cmd != nil {
		cmd.Start()
		cmd.Wait()
	}

	if cmd := t.ActiveLesson.setup(); cmd != nil {
		cmd.Start()
		cmd.Wait()
	}

	if t.ActiveLesson.Example != "" {
		t.generateExample()
	}

	var answer = false

	for answer == false {
		t.ActiveLesson.teach()

		cmd, err := prompt(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err.Error())
		}

		answer = t.checkAnswer(cmd)

		out, err := exec.Command("/bin/sh", "-c", cmd).CombinedOutput()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err.Error())
		}
		fmt.Printf("%s", out)

		if answer {
			t.success()
			return
		} else {
			t.failure()
		}
	}

}

// CheckAnswer splits the command on a newline and checks the answer
func (t *Tutorial) checkAnswer(cmd string) bool {
	answer := bytes.Split([]byte(cmd), []byte("\n"))

	if bytes.Equal(answer[0], []byte(t.ActiveLesson.Answer)) {
		return true
	}
	return false
}

// Setup provisions lesson resources
func (l *Lesson) setup() *exec.Cmd {
	if len(l.Setup) > 0 {
		for _, cmd := range l.Setup {
			return exec.Command("/bin/sh", "-c", cmd)
		}
	}
	return nil
}

// Teardown tears down lesson resources
func (l *Lesson) teardown() *exec.Cmd {
	if len(l.Teardown) > 0 {
		for _, cmd := range l.Teardown {
			return exec.Command("/bin/sh", "-c", cmd)
		}
	}
	return nil
}

func (t *Tutorial) generateExample() {
	examplesDir := fmt.Sprintf("%s%s", t.Directory, t.ActiveLesson.Example)
	if err := os.MkdirAll(examplesDir, 0700); err != nil {
		log.Fatalf("Failed to create directories: %s", err.Error())
	}

	templates := fmt.Sprintf("%s%s", templatePath(), t.ActiveLesson.Example)
	files, _ := ioutil.ReadDir(templates)

	for _, file := range files {
		sourcePath := fmt.Sprintf("%s/%s",templates, file.Name())
		inputFile, err := os.Open(sourcePath)
		if err != nil {
			log.Fatalf("Failed to open the file:%s", err)
		}

		destPath := fmt.Sprintf("%s/%s",examplesDir,file.Name())
		outputFile, err := os.Create(destPath)
		if err != nil {
			log.Fatalf("Failed to create the file:%s", err)
		}

		io.Copy(outputFile, inputFile)
	}
}

// Teach returns a lesson exercise
func (l *Lesson) teach() {
	fmt.Fprintln(os.Stdout, l.Title)
	fmt.Fprintln(os.Stdout, lbreak())
	fmt.Fprintln(os.Stdout, l.Exercise)
}

// Explain returns a lesson explanation
func (l *Lesson) explain() {
	fmt.Fprintln(os.Stdout, l.Explanation)
}

// Failure represents a lesson failed
func (t *Tutorial) failure() {
	fmt.Println()
	fmt.Fprintln(os.Stdout, "Command was not correct.")
	fmt.Println()
}

// Success represents a lesson succeeded
func (t *Tutorial) success() {
	fmt.Println()
	fmt.Fprintln(os.Stdout, "Correct!")
	fmt.Println()

	lessLen := len(t.Lessons)

	if t.ActiveLessonId == lessLen-1 {
		t.reset()
		cat := catMap[t.Category]
		capitalize := bytes.Title([]byte(Categories[cat]))
		msg := fmt.Sprintf("%s Tutorial Complete!", capitalize)
		fmt.Fprintln(os.Stdout, msg)
		return
	}
	t.ActiveLesson = t.Lessons[t.ActiveLessonId+1]
	t.ActiveLessonId = t.ActiveLessonId + 1
	t.save()
	// Write to Tutorial Json
	t.NextLesson()
}

func (t *Tutorial) save() {
	tuts := t.Tutorials
	cat := catMap[t.Category]
	tuts[cat].ActiveLessonId = t.ActiveLessonId
	tuts[catMap["docker"]].Directory = t.Directory
	tuts[catMap["docker-compose"]].Directory = t.Directory
	tuts[catMap["swarm"]].Directory = t.Directory

	json, err := json.Marshal(tuts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
	}

	tutsConfig, _ := ConfigFiles(t.Category)
	ioutil.WriteFile(tutsConfig, json, 0644)
}

// Reset resources and lesson progress
func (t *Tutorial) reset() {
	tuts := t.Tutorials
	cat := catMap[t.Category]
	tuts[cat].ActiveLessonId = 0

	json, err := json.Marshal(tuts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
	}

	tutsConfig, _ := ConfigFiles(t.Category)
	ioutil.WriteFile(tutsConfig, json, 0644)
}

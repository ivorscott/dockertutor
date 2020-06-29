// Tutor provides state management for tutorials and their lessons
package tutor

//go:generate binclude

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/lu4p/binclude"
)

// Tutorial manages active tutorial state
type Tutorial struct {
	Config            `json:"-"`
	Category          string
	ActiveLessonIndex int
	ActiveLesson      Lesson `json:"-"`
	Lessons           `json:"-"`
}

// Config stores application Tutorials and the initialized practice directory.
// Config is used to retrieve persisted data found in tutor_config.json.
type Config struct {
	Directory string
	Tutorials []Tutorial
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

var ConfigFile = "tutor_config.json"
var CommandNotFoundExitCode = 127

// prompt prompts the user for input and reads standard in
func prompt(stdin io.Reader) (string, error) {
	fmt.Print("\n\n> ")
	reader := bufio.NewReader(stdin)
	return reader.ReadString('\n')
}

// configFilePath returns the path to the tutor configuration on the user's local filesystem
func configFilePath(path string) string {
	return fmt.Sprintf("%s/%s", path, ConfigFile)
}

// NewConfig creates a new tutor configuration file on the user's local filesystem
func NewConfig(path string) error {
	conf := configFilePath(path)
	data := fmt.Sprintf(`{ "Directory":"%s"}`, path)
	if err := ioutil.WriteFile(conf, []byte(data), 0700); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(os.Stdout, "directory initialized"); err != nil {
		return err
	}
	return nil
}

// OpenOrCreateConfig opens the tutor configuration if it exists, otherwise it creates it
func OpenOrCreateConfig(path string) error {
	f, err := OpenConfig()
	if err != nil {
		return NewConfig(path)
	}
	defer f.Close()

	if _, err := fmt.Fprintln(os.Stdout, "folder already initialized"); err != nil {
		return err
	}
	return nil
}

// OpenConfig opens the tutor configuration by checking in 3 places:
// the current directory, the parent directory and the grandparent directory.
// If not found in one location it function moves up a directory.
func OpenConfig() (*os.File, error) {
	var f *os.File
	var err error
	conf := ConfigFile
	ctx := ""

	for tries := 0; tries < 3; tries++ {
		f, err = func() (*os.File, error) {
			f, err := os.Open(fmt.Sprintf("%s%s", ctx, conf))
			if err != nil {
				defer f.Close()
				return nil, err
			}
			return f, nil
		}()

		if err == nil {
			break
		}

		ctx += "../"

		if tries == 2 {
			if err != nil {
				return nil, fmt.Errorf("Configuration missing. Directory is not initialized. \n\n" +
					"Try running: dockertutor init or dockertutor help\n\n")
			}
		}
	}

	return f, nil
}

// NewApp returns a new tutor application for category
func NewTutorial(f *os.File, lessonData []byte, category string) (*Tutorial, error) {
	conf := &Config{}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, conf); err != nil {
		return nil, err
	}

	if len(conf.Tutorials) == 0 {
		tuts := &[]Tutorial{}
		tutsData := `[
			{
				"Category": "docker",
				"ActiveLessonIndex": 0
			},
			{
				"Category": "docker-compose",
				"ActiveLessonIndex": 0
			},
			{
				"Category": "swarm",
				"ActiveLessonIndex": 0
			}
		]`
		if err := json.Unmarshal([]byte(tutsData), tuts); err != nil {
			return nil, err
		}
		conf.Tutorials = *tuts

		j, err := json.Marshal(conf)
		if err != nil {
			return nil, err
		}

		if err := ioutil.WriteFile(ConfigFile, j, 0644); err != nil {
			return nil, err
		}
	}

	l, err := NewLessons(lessonData)
	if err != nil {
		return nil, err
	}

	cat := catMap[category]
	lessons := *l
	lessonIndex := conf.Tutorials[cat].ActiveLessonIndex

	tut := &Tutorial{
		Config:            *conf,
		Category:          category,
		ActiveLessonIndex: lessonIndex,
		ActiveLesson:      lessons[lessonIndex],
		Lessons:           lessons,
	}

	return tut, nil
}

// Tutorial returns a lesson exercise
func (t *Tutorial) Welcome() error {
	if _, err := fmt.Fprintln(os.Stdout, IntroMap[t.Category]); err != nil {
		return err
	}
	return nil
}

// Next fetches the next lesson
func (t *Tutorial) NextLesson() error {
	if cmd := t.ActiveLesson.teardown(); cmd != nil {
		if err := cmd.Start(); err != nil {
			return err
		}
		if err := cmd.Wait(); err != nil {
			return err
		}
	}

	if cmd := t.ActiveLesson.setup(); cmd != nil {
		if err := cmd.Start(); err != nil {
			return err
		}
		if err := cmd.Wait(); err != nil {
			return err
		}
	}

	if t.ActiveLesson.Example != "" {
		if err := t.generateExample(); err != nil {
			return err
		}
	}

	var answer = false
	for answer == false {
		if err := t.ActiveLesson.teach(); err != nil {
			return err
		}

		cmd, err := prompt(os.Stdin)
		if err != nil {
			return err
		}

		answer = t.checkAnswer(cmd)
		command := exec.Command("/bin/sh", "-c", cmd)

		out, err := command.CombinedOutput()
		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				if exitError.ExitCode() != CommandNotFoundExitCode {
					return err
				}
			}
		}
		fmt.Printf("%s", out)

		if answer {
			return t.success()
		}
		if err := t.failure(); err != nil {
			return err
		}
	}
	return nil
}

// checkAnswer splits the command on a newline and checks the answer
func (t *Tutorial) checkAnswer(cmd string) bool {
	answer := bytes.Split([]byte(cmd), []byte("\n"))
	if bytes.Equal(answer[0], []byte(t.ActiveLesson.Answer)) {
		return true
	}
	return false
}

// generateExample copies a statik example to the user's directory
func (t *Tutorial) generateExample() error {
	binclude.Include("../examples")

	// create example folder structure in user's local directory
	usd := fmt.Sprintf("%s/%s", t.Directory, t.ActiveLesson.Example)
	if err := os.MkdirAll(usd, 0700); err != nil {
		return err
	}

	// copy static examples and move them to the user's local directory
	exd := fmt.Sprintf("%s/%s", "../examples", t.ActiveLesson.Example)
	files, err := BinFS.ReadDir(exd)
	if err != nil {
		return err
	}

	for _, file := range files {
		src := fmt.Sprintf("%s/%s", exd, file.Name())
		inf, err := BinFS.Open(src)
		if err != nil {
			return err
		}

		dest := fmt.Sprintf("%s/%s", usd, file.Name())
		outf, err := os.Create(dest)
		if err != nil {
			return err
		}

		if _, err := io.Copy(outf, inf); err != nil {
			return err
		}
	}
	return nil
}

// failure represents a lesson failed
func (t *Tutorial) failure() error {
	if _, err := fmt.Fprint(os.Stdout, "\nCommand was not correct.\n"); err != nil {
		return err
	}
	return nil
}

// success represents a lesson succeeded
func (t *Tutorial) success() error {
	if _, err := fmt.Fprint(os.Stdout, "\nCorrect!\n"); err != nil {
		return err
	}

	lessLen := len(t.Lessons)
	if t.ActiveLessonIndex == lessLen-1 {
		if err := t.reset(); err != nil {
			return err
		}
		cat := catMap[t.Category]
		category := bytes.Title([]byte(Categories[cat]))
		msg := "Tutorial Complete!"
		if _, err := fmt.Fprintln(os.Stdout, fmt.Sprintf("%s %s", category, msg)); err != nil {
			return err
		}
		return nil
	}

	t.ActiveLesson = t.Lessons[t.ActiveLessonIndex+1]
	t.ActiveLessonIndex = t.ActiveLessonIndex + 1
	if err := t.save(); err != nil {
		return err
	}
	return t.NextLesson()
}

// reset resources and lesson progress
func (t *Tutorial) reset() error {
	cat := catMap[t.Category]
	t.Config.Tutorials[cat].ActiveLessonIndex = 0

	j, err := json.Marshal(t.Config)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(configFilePath(t.Config.Directory), j, 0644); err != nil {
		return err
	}
	return nil
}

// save updates persistent storage configuration
func (t *Tutorial) save() error {
	cat := catMap[t.Category]
	t.Config.Tutorials[cat].ActiveLessonIndex = t.ActiveLessonIndex

	j, err := json.Marshal(t.Config)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(configFilePath(t.Config.Directory), j, 0644); err != nil {
		return nil
	}
	return nil
}

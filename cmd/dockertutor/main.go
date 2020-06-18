package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/ivorscott/dockertutor/lessons"
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

	tut := flag.String("c", lessons.Categories[0], "Select tutorial category")
	flag.Parse()

	ll := &lessons.Lessons{}
	filename := fmt.Sprintf("lessons/%s.%s", *tut, "json")

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}

	json.Unmarshal(file, ll)
	fmt.Printf("%v", ll)

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

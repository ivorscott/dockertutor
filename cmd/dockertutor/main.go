package main

//go:generate binclude

import (
	"fmt"
	"github.com/ivorscott/dockertutor/tutor"
	"github.com/lu4p/binclude"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.SetFlags(0)
		log.Fatal(err)
	}
}

func run() error {
	binclude.Include("../../lessons")

	app := &cli.App{
		Name:      "dockertutor",
		Usage:     "Interactive tutorial for learning docker",
		UsageText: "dockertutor [OPTIONS] [COMMAND]",
		Commands: []*cli.Command{
			{
				Name:  "init",
				Usage: "initialize practice directory",
				Action: func(c *cli.Context) error {
					dirname := c.Args().First()
					path, err := os.Getwd()
					if err != nil {
						return err
					}

					if dirname != "" {
						dirPath := fmt.Sprintf("%s/%s", path, dirname)
						if _, err := ioutil.ReadDir(dirPath); err != nil {
							if _, err := fmt.Fprintf(os.Stdout, "directory does not exist. "+
								"creating (%s) directory...\n", dirname); err != nil {
								return err
							}
							if err := os.Mkdir(dirPath, 0700); err != nil {
								return err
							}
							return tutor.NewConfig(dirPath)
						}
						return tutor.OpenOrCreateConfig(dirPath)
					}
					return tutor.OpenOrCreateConfig(path)
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "category",
				Aliases:     []string{"c"},
				Usage:       "choices: docker|docker-compose|swarm",
				Value:       "docker",
				DefaultText: "docker",
			},
		},
		Action: func(c *cli.Context) error {

			f, err := tutor.OpenConfig()
			if err != nil {
				return err
			}
			defer f.Close()

			config, err := ioutil.ReadAll(f)
			if err != nil {
				return err
			}

			cFile := fmt.Sprintf("../../lessons/%s.json", c.String("category"))
			bf, err := BinFS.Open(cFile)
			if err != nil {
				return err
			}
			defer bf.Close()

			lessons, err := ioutil.ReadAll(bf)
			if err != nil {
				return err
			}

			t, err := tutor.NewTutorial(config, lessons, c.String("category"))
			if err != nil {
				return err
			}

			if t.ActiveLessonIndex == 0 {
				if err := t.Welcome(); err != nil {
					return err
				}
			}

			if err := t.NextLesson(); err != nil {
				return err
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		return err
	}

	return nil
}

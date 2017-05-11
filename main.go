package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/cli"
)

type Root struct {
	root     string
	troot    string
	have     string
	finished string
}

var root = NewRoot()

func (r Root) GetRootTasks() string {
	return r.troot
}

func (r Root) GetListTasks() []string {
	return r.GetList(r.GetRootTasks())
}
func (r Root) GetRootHave() string {
	return r.have
}

func (r Root) GetListHave() []string {
	return r.GetList(r.GetRootHave())
}
func (r Root) GetList(rpath string) []string {
	var result []string
	err := filepath.Walk(rpath,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			rel, err := filepath.Rel(rpath, path)
			result = append(result, rel)
			return nil
		})
	if err != nil {
		fmt.Println(err)
	}
	return result
}

func main() {
	c := cli.NewCLI("app", "1.0.0")

	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"init": func() (cli.Command, error) {
			return &Init{}, nil
		},
		"list": func() (cli.Command, error) {
			return &List{}, nil
		},
		"add": func() (cli.Command, error) {
			return &Add{}, nil
		},
		"show": func() (cli.Command, error) {
			return &Show{}, nil
		},
		"done": func() (cli.Command, error) {
			return &Done{}, nil
		},
		"edit": func() (cli.Command, error) {
			return &Edit{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)
}

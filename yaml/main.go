package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// Test is used to for yaml loading
type Test struct {
	Name               string `yaml:"name"`
	Version            string
	Executable         string `yaml:"exec"`
	WorkingDirectory   string
	DefaultParameters  []string
	DefaultEnvironment []string
	ConnectionDetails  string
	Modes              []string
}

var bytes = `
# Comment to ignore
name: test
exec: "%ENV%\file"
`

func main() {
	fmt.Println("Start")

	t := new(Test)
	err := yaml.Unmarshal([]byte(bytes), t)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Parsed is %v+\n", t)

	t.Name = "Example"
	t.Version = "1.0.0"
	t.WorkingDirectory = ""
	t.DefaultParameters = []string{"-c", "file name"}
	t.DefaultEnvironment = []string{"NAME=VALUE"}
	t.ConnectionDetails = "localhost:8080/path"
	t.Modes = []string{"auto", "aware", "single"}

	out, err := yaml.Marshal(t)

	fmt.Printf("Example is:\n%s\n", string(out))
}

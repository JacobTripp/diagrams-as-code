/*
genDiagram creates the boiler plate code when defining a new diagram type. It
will create a new child directory in $PWD if there are no name collisions
already, if there is a name collision then the command will panic.

Usage:

	genDiagram [-d yaml]

The flags are:

	-d
			Diagram definition path. The definition of the diagram that you want to
			create.

*/
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/JacobTripp/diagrams-as-code/tools/builder"
)

var yml = flag.String(
	"d",
	"",
	"Path to a yaml file containing a (d)iagram definition",
)

func readYaml(p string) ([]byte, error) {
	data, err := os.ReadFile(p)
	if err != nil {
		return nil, (err)
	}
	return data, nil
}

func getDirs(pwd string) (map[string]struct{}, error) {
	files, err := os.ReadDir(pwd)
	if err != nil {
		return nil, err
	}

	dirIndex := make(map[string]struct{}, len(files))

	for _, file := range files {
		if file.IsDir() {
			dirIndex[file.Name()] = struct{}{}
		}
	}
	return dirIndex, nil
}

func parse(data []byte) (builder.YamlData, error) {
	rt, err := builder.ImportYaml(data)
	if err != nil {
		return rt, err
	}
	return rt, nil
}

func makeDir(name string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	dirs, err := getDirs(pwd)
	if err != nil {
		return err
	}
	if _, ok := dirs[name]; ok {
		return fmt.Errorf("the diagram name '%s' already exists!", name)
	}
	err = os.Mkdir(name, 0777)
	if err != nil {
		return err
	}
	return nil
}

func readAndParse(n string) (builder.YamlData, error) {
	data, err := readYaml(n)
	if err != nil {
		r := new(builder.YamlData)
		return *r, err
	}

	rt, err := parse(data)
	if err != nil {
		return rt, err
	}
	return rt, nil
}

func main() {
	flag.Parse()

	diagram, err := readAndParse(*yml)
	if err != nil {
		panic(err)
	}

	dirName := strings.ToLower(diagram.Diagram.Name)
	err = makeDir(dirName)
	if err != nil {
		panic(err)
	}
	packagePath := fmt.Sprintf(
		"%s/%s.go",
		dirName,
		strings.ToLower(diagram.Diagram.Name),
	)
	packageFile, err := os.Create(packagePath)
	if err != nil {
		panic(err)
	}
	builder.WritePackage(packageFile, diagram)
}

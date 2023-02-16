package main

import (
	"bytes"
	"crypto/md5"
	"flag"
	"fmt"
	"go/format"
	"html/template"
	"os"
	"path"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

var tmplPath = flag.String(
	"t",
	"./diagram.go.tmpl",
	"path to the template to use for generating code",
)

var outPath = flag.String(
	"o",
	"",
	"the location where you want to create the new directory",
)

var yamlData = flag.String(
	"d",
	"",
	"The yaml file with the diagram definition",
)

func recieverName(name string) string {
	name = strings.ReplaceAll(name, " ", "")
	if name == "" {
		panic("empty reciever name")
	}
	getTo := 2
	if len(name) < getTo {
		getTo = len(name)
	}
	return fmt.Sprintf("_%s", strings.ToLower(name[:getTo]))
}

func renderedAt() int64 {
	return time.Now().Unix()
}

func loadTemplate(tmplPath string) (*template.Template, string) {
	fnMap := template.FuncMap{
		"recieverName": recieverName,
		"renderedAt":   renderedAt,
		"yamlHash":     func() string { return "yaml hash" },
		"templateHash": func() string { return "template hash" },
		"packageName":  func(d graph) string { return strings.ReplaceAll(d.Name, " ", "") },
		"graphType":    func(d graph) string { return d.Type },
		"getNodeName":  func(d node) string { return d.Name },
		"getEdgeName":  func(d edge) string { return d.Name },
		//"escapeChars":  func(d string) string { return strings.ReplaceAll(d, "\"", "\\\"") },
	}
	name := path.Base(tmplPath)
	tmpl, err := template.New(name).Funcs(fnMap).ParseFiles(tmplPath)
	if err != nil {
		panic(err)
	}
	// I know, crappy to read the file twice
	f, err := os.ReadFile(tmplPath)
	if err != nil {
		panic(err)
	}
	return tmpl, fmt.Sprintf("%x", md5.Sum(f))
}

type entity struct {
	Name        string    `yaml:"name"`
	Description string    `yaml:"description"`
	Attributes  attribute `yaml:"attributes"`
}
type node struct {
	entity `yaml:",inline"`
}
type edge struct {
	entity   `yaml:",inline"`
	Disallow rule `yaml:"disallow"`
}
type attribute map[string]string
type rule struct {
	To   []string `yaml:"to"`
	From []string `yaml:"From"`
}

type graph struct {
	entity `yaml:",inline"`
	Type   string `yaml:"type"`
}

type data struct {
	Diagram graph  `yaml:"diagram"`
	Nodes   []node `yaml:"nodes"`
	Edges   []edge `yaml:"edges"`
}

func loadYaml(path string) (*data, string) {
	d := &data{}
	y, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(y, d)
	if err != nil {
		panic(err)
	}
	return d, fmt.Sprintf("%x", md5.Sum(y))
}

// remove this, put this in the write function
func openOut(templateSum, yamlSum, outPath string) *os.File {
	f, err := os.OpenFile(outPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	return f
}

func render(t *template.Template, data *data) string {
	var out = bytes.NewBuffer([]byte{})
	err := t.Execute(out, data)
	if err != nil {
		panic(err)
	}
	fmted, err := format.Source(out.Bytes())
	if err != nil {
		panic(fmt.Sprintf("format error: %s", err))
	}
	return fmt.Sprintf("%s", fmted)
}

package newDiagram

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"go/format"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
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

func yamlHash() string {
	return "yaml hash"
}
func templateHash() string {
	return "template hash"
}

func loadTemplate(tpath string) (*template.Template, string) {
	fnMap := template.FuncMap{
		"recieverName": recieverName,
		"renderedAt":   renderedAt,
		"yamlHash":     yamlHash,
		"templateHash": templateHash,
		"packageName":  func(d string) string { return strings.ReplaceAll(d, " ", "") },
		"graphType":    func(d graph) string { return d.Type },
		"getNodeName":  func(d node) string { return d.Name },
		"getEdgeName":  func(d edge) string { return d.Name },
		//"escapeChars":  func(d string) string { return strings.ReplaceAll(d, "\"", "\\\"") },
	}
	name := filepath.Base(tpath)
	tmpl, err := template.New(name).Funcs(fnMap).ParseFiles(tpath)
	if err != nil {
		panic(err)
	}
	// I know, crappy to read the file twice
	f, err := os.ReadFile(tpath)
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
	entity
	Type string
}

type Diagram interface {
	GetString(string) string
	Get(string) interface{}
}

func Render(out io.Writer, templatePath string, data Diagram) error {
	t, _ := loadTemplate(templatePath)
	var unfmted = bytes.NewBuffer([]byte{})
	err := t.Execute(unfmted, data)
	if err != nil {
		return err
	}
	fmted, err := format.Source(unfmted.Bytes())
	if err != nil {
		return fmt.Errorf("%w: format error", err)
	}
	fmt.Fprintf(out, "%s", fmted)
	return nil
}

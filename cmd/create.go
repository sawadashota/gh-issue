package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"log"
	"path"
	"regexp"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
	"reflect"
	"github.com/sawadashota/gh-issue/issue"
)

type issueYaml struct {
	input string
	abs   string
	body map[string][]map[string]interface{}
}

var i *issueYaml

var Create = &cobra.Command{
	Use:   "create -f [filepath]",
	Short: "Create issue at GitHub",
	Long:  `Create issue at GitHub`,
	Args:  cobra.MaximumNArgs(0),
	PreRun: func(cmd *cobra.Command, args []string) {
		i = newYaml(file)
		if !i.isYamlExtension() {
			log.Fatal("File extension should be issueYaml or yml")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		issues := issue.New("aaa") // todo envchainからgh tokenを参照する
		for _, i := range i.body["issues"] {
			var ops []issue.Option

			if contains(i, "assignee") {
				ops = append(ops, issue.WithAssignee(i["assignee"].(string)))
			}

			if contains(i, "labels") {
				ops = append(ops, issue.WithLabels(getLabels(i)))
			}

			issues.AddIssue(i["title"].(string), ops...)
		}
		fmt.Printf("%v\n", issues) // FIXME 消す
	},
}

func getLabels(i map[string]interface{}) []issue.Label {
	slice := reflect.ValueOf(i["labels"])

	if slice.Kind() != reflect.Slice {
		log.Fatal("labels should be list.")
	}

	var labels []issue.Label
	for k := 0; k < slice.Len(); k++ {
		var label issue.Label
		label.Name = slice.Index(k).Interface().(string)
		labels = append(labels, label)
	}

	return labels
}

// Set absolution path for issueYaml file
func newYaml(file string) *issueYaml {
	i := &issueYaml{input: file}

	pwd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	i.abs = path.Join(pwd, i.input)
	i.read()

	return i
}

func (i *issueYaml) read() {
	buf, err := ioutil.ReadFile(i.abs)

	if err != nil {
		log.Fatal(err)
	}

	m := make(map[string][]map[string]interface{})
	err = yaml.Unmarshal(buf, &m)

	if err != nil {
		log.Fatal(err)
	}

	i.body = m
}

func (i *issueYaml) isYamlExtension() bool {
	r := regexp.MustCompile(`\.(yaml|yml)$`)
	return r.MatchString(i.abs)
}

func contains(s map[string]interface{}, e string) bool {
	for key := range s {
		if e == key {
			return true
		}
	}
	return false
}

package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"

	"github.com/sawadashota/gh-issue"
	"github.com/sawadashota/gh-issue/eloquent"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type issueYaml struct {
	input string
	abs   string
	body  map[string]interface{}
}

var i *issueYaml

var Create = &cobra.Command{
	Use:   "create -f [filepath] -r [repository]",
	Short: "Create issue at GitHub",
	Long:  `Create issue at GitHub`,
	Args:  cobra.MaximumNArgs(0),
	PreRun: func(cmd *cobra.Command, args []string) {
		i = newYaml(file)
		if !i.isYamlExtension() {
			log.Fatal("File extension should be issueYaml or yml")
		}

		if repo == "" {
			log.Fatal("[-r --repo] Repository should be present")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		token, err := getToken()
		if err != nil {
			log.Fatal(err)
		}

		o, r, err := splitOwnerRepo(repo)

		if err != nil {
			log.Fatal(err)
		}

		issues := issues(o, r, token, i.body["issues"].([]interface{}))
		results := issues.Create()

		stdoutAllError(results)
	},
}

func splitOwnerRepo(repo string) (owner string, name string, err error) {
	arr := strings.Split(repo, "/")
	if len(arr) != 2 {
		return "", "", fmt.Errorf("invalid repository name: %s", repo)
	}

	owner = arr[0]
	name = arr[1]
	return
}

func stdoutAllError(results *[]ghissue.Result) {
	errCount := 0
	for _, result := range *results {
		if result.HasError() {
			errCount++

			if errCount == 1 {
				eloquent.NewError("\n************* Error List *************\n\n").Important().Exec()
			}

			eloquent.NewError("%d. %v\n%v\n\n", errCount, result.Issue.Title, result.Error.Error()).Exec()
		}
	}

	eloquent.NewError("%d errors occurred.\n", errCount).Important().Exec()
}

// Yamlのissues以下を受け取り、構造体を返す
func issues(owner string, repo string, token string, yaml []interface{}) *ghissue.Issues {
	issues := ghissue.New(owner, repo, token)
	for _, issue := range yaml {
		var ops []ghissue.Option

		title, err := getString(issue, "title")
		if err != nil {
			log.Fatal(err)
		}

		assignee, err := getString(issue, "assignee")
		if err == nil {
			ops = append(ops, ghissue.WithAssignee(assignee))
		}

		body, err := getString(issue, "body")
		if err == nil {
			ops = append(ops, ghissue.WithBody(body))
		}

		labels, err := getSlice(issue, "labels")

		if err == nil {
			for _, label := range labels {
				ops = append(ops, ghissue.WithLabel(label))
			}
		}

		issues.AddIssue(title, ops...)
	}

	return issues
}

// Yamlから値がstringの値を取り出す
func getString(yaml interface{}, key string) (string, error) {
	for k, v := range yaml.(map[interface{}]interface{}) {
		if fmt.Sprintf("%v", k) == key {
			return fmt.Sprintf("%v", v), nil
		}
	}

	return "", fmt.Errorf("key: \"%v\" is not exist in yaml file", key)
}

// Yamlから値が[]stringの値を取り出す
func getSlice(yaml interface{}, key string) ([]string, error) {
	var slice []string
	for k, v := range yaml.(map[interface{}]interface{}) {
		if fmt.Sprintf("%v", k) == key {
			for _, label := range v.([]interface{}) {
				slice = append(slice, fmt.Sprintf("%v", label))
			}
			return slice, nil
		}
	}

	return slice, fmt.Errorf("key: %v is not exist", key)
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

	m := make(map[string]interface{})
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

func contains(s interface{}, e string) bool {
	for key := range s.(map[string]interface{}) {
		if e == key {
			return true
		}
	}
	return false
}

func getToken() (string, error) {
	if !executable("envchain") {
		eloquent.NewWarning("Please install envchain", "https://github.com/sorah/envchain").Exec()
		log.Fatal("Command envchain is not exists.")
	}

	command := exec.Command("envchain", "gh-issue", "env")
	res, err := command.Output()

	if err != nil {
		return "", err
	}

	r := regexp.MustCompile(EnvchainEnv + `=(.+)`)
	matches := r.FindStringSubmatch(string(res))

	if len(matches) < 2 {
		return "", fmt.Errorf("Cannot Find GITHUB_TOKEN\n")
	}

	return matches[1], nil
}

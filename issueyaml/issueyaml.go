package issueyaml

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/sawadashota/gh-issue"
	"gopkg.in/yaml.v2"
)

type Yaml struct {
	input string
	abs   string
	body  map[string]interface{}
}

var yamlExtRegexp *regexp.Regexp

func init() {
	yamlExtRegexp = regexp.MustCompile(`\.(yaml|yml)$`)
}

func (y *Yaml) OwnerRepo() (owner string, name string, err error) {
	meta := y.body["meta"].(interface{})
	repo, err := getString(meta, "repo")
	if err != nil {
		return "", "", err
	}

	r, o, err := splitOwnerRepo(repo)
	if err != nil {
		return "", "", err
	}

	return o, r, nil
}

// Yamlのissues以下を受け取り、構造体を返す
func (y *Yaml) Issues(token string) (*ghissue.Issues, error) {
	repo, owner, err := y.OwnerRepo()
	if err != nil {
		return nil, err
	}

	issues := ghissue.NewIssues(owner, repo, token)
	for _, issue := range y.issues() {
		var ops []ghissue.IssueOption

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

	return issues, nil
}

// Set absolution path for issueYaml file
func New(fp string) (*Yaml, error) {
	if !isYamlExtension(fp) {
		return nil, fmt.Errorf("file extension should be issueYaml or yml")
	}

	i := &Yaml{input: fp}

	pwd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	if strings.HasPrefix(fp, "/") {
		i.abs = fp
	} else {
		i.abs = path.Join(pwd, i.input)
	}
	i.read()

	return i, nil
}

func (y *Yaml) read() {
	buf, err := ioutil.ReadFile(y.abs)

	if err != nil {
		log.Fatal(err)
	}

	m := make(map[string]interface{})
	err = yaml.Unmarshal(buf, &m)

	if err != nil {
		log.Fatal(err)
	}

	y.body = m
}

func (y *Yaml) issues() []interface{} {
	return y.body["issues"].([]interface{})
}

// yamlの拡張子かどうか
func isYamlExtension(fp string) bool {
	return yamlExtRegexp.MatchString(fp)
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

func splitOwnerRepo(repo string) (owner string, name string, err error) {
	arr := strings.Split(repo, "/")
	if len(arr) != 2 {
		return "", "", fmt.Errorf("invalid repository name: %s", repo)
	}

	owner = arr[0]
	name = arr[1]
	return
}

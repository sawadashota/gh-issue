gh-issue
===

Create GitHub issues from YAML.

Required
---

* [envchain](https://github.com/sorah/envchain)

Usage
---

### Set GitHub token

```bash
gh-issue set -t [GitHub token]
```

Then, you will be required type password for access keychain.

### Create `issues.yml`

You can make template yaml file by `gh-issue init`

```yml
meta:
  repo: sawadashota/gh-issue

issues:
  - title: issue title 1
    assignee: sawadashota
    body: |-
      Example title
      ===
      Example body

      * Example list
    labels:
      - enhancement
      - bug
  - title: issue title 2
```

### Create issues on GitHub

```bash
gh-issue create -f [path to yaml]
```

For example

```bash
gh-issue create -f ./issues.yml
```

Installation
---

```bash
brew tap sawadashota/homebrew-cheers
brew install gh-issue
```

or

```bash
go get -u github.com/sawadashota/gh-issue
```

License
---

MIT

Author
---

Shota Sawada

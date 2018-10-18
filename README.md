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
$ gh-issue set -t [GitHub token]
```

Then, you will be required type password for access keychain.

### Create Issues

```bash
$ gh-issue
```


For example like this.

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

Saved and quit editor, then start creating issues on GitHub!

### Edit config

```bash
$ gh-issue edit
```

```toml
editor = "vim"
template = """# template
meta:
  repo: owner/reponame

issues:
  - title: issue title 1
    assignee: assignee
    body: ""
    labels:
      - enhancement"""
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

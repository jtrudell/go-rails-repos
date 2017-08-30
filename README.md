# go-rails-repos
A Go CLI that updates local rails repositories from Github and bundles, drops, creates, migrates and seeds databases

## Installation
```bash
go get githb.com/jtrudell/go-rails-repos
```

## Usage
```bash
$GOPATH/bin/go-rails-repos -repos <list your git repos here without commas>
```

Use with one or more multiple git repositories in the same directory (e.g. your "projects" folder). Does not support multiple directories.

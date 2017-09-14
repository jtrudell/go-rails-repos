# go-rails-repos
A Go CLI that updates your local rails repositories from origin/master, bundles and then runs the rake tasks to drop, create, migrate and seed the repo database.

## Installation
```bash
go get githb.com/jtrudell/go-rails-repos
```

## Usage
```bash
$GOPATH/bin/go-rails-repos -repos <list your git repos here without commas>
```

Use with one or more git repositories in the same directory (e.g. your "projects" folder). Does not support multiple directories.

Runs the following commands in each repository:

```bash
git stash
git checkout master
git pull origin master
bundle install
bundle exec rake db:drop
bundle exec rake db:create
bundle exec rake db:migrate
bundle exec rake db:seed
```

## Why?
Because at work we have a bash script that does this but it is slow. This is faster, because of concurrency (and, depending on your computer, parallelism using GOMAXPROCS). Also, Go is fun.

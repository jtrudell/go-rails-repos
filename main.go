package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os/exec"
	"runtime"
	"time"
)

// replace with your default repos here, which will all need to be in the same directory (e.g., "projects")
// repos can be changed at runtime with with flag -repos <list here>
var repos = []string{
	"",
	"",
	"",
}

// counter initialized to manage runRepo go routines and ensure main doesn't exit until complete
var counter int = 0

// gitCommands takes the dir string and runs gitStash, gitCheckoutMaster and gitPullMaster for it
func gitCommands(dir string) {
	gitStash(dir)
	gitCheckoutMaster(dir)
	gitPullMaster(dir)
}

// gitStash takes the dir string and executes git stash in the directory
func gitStash(dir string) {
	cmd := exec.Command("git", "stash")
	execute(cmd, dir)
	fmt.Printf("Stashing changes on %s...\n", dir)
}

// gitCheckoutMaster takes the dir string and executes git checkout master in the directory
func gitCheckoutMaster(dir string) {
	cmd := exec.Command("git", "checkout", "master")
	execute(cmd, dir)
	fmt.Printf("Checking out master on %s...\n", dir)
}

// gitPullMaster takes the dir string and executes git pull origin master in the directory
func gitPullMaster(dir string) {
	cmd := exec.Command("git", "pull", "origin", "master")
	execute(cmd, dir)
	fmt.Printf("Pulling origin master on %s...\n", dir)
}

// bundle takes the directory string and executes bundle install in the directory
func bundle(dir string) {
	cmd := exec.Command("bundle", "install")
	execute(cmd, dir)
	fmt.Printf("Bundling %s...\n", dir)
}

// rakeTasks takes the dir string and runs various rake tasks
func rakeTasks(dir string) {
	rake(dir, "db:drop")
	rake(dir, "db:create")
	rake(dir, "db:migrate")
	rake(dir, "db:seed")
}

// rake takes the dir string and the task string and executes the rake task in the directory
func rake(dir string, task string) {
	cmd := exec.Command("bundle", "exec", "rake", task)
	execute(cmd, dir)
	fmt.Printf("Running bundle exec rake %s on %s...\n", task, dir)
}

// execute takes a pointer to *exec.Cmd and dir string and sets the correct directory
// it runs the command and logs any error
func execute(cmd *exec.Cmd, dir string) {
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		log.Fatalln(color.RedString("Something broke on "+dir+":"), err)
	}
}

// runRepo takes the dir string and runs gitCommands, bundle and rakeTasks for it
// It also benchmarks the go routine run time
func runRepo(dir string) {
	start := time.Now()
	blue := color.New(color.FgBlue)

	gitCommands(dir)
	bundle(dir)
	rakeTasks(dir)

	duration := time.Since(start)
	blue.Printf("Done with %s. Time elapsed: %v\n", dir, duration)
	counter++
}

// reposFlag initializes -repos flag and checks that repos has been populated
func reposFlag() {
	repoPtr := flag.Bool("repos", false, "a list of repositories separated by spaces")
	flag.Parse()

	if *repoPtr {
		repos = flag.Args()
	}

	if len(flag.Args()) == 0 || repos[0] == "" {
		err := errors.New("You must pass in a list of repos using the -repos flag")
		log.Fatalln(err)
	}
}

// init sets GOMAXPROCS and calls repos flag
func init() {
	runtime.GOMAXPROCS(4)
	reposFlag()
}

// main iterates through each repo and calls a go routine for runRepo(repo)
// It also benchmarks the runtime of the program
func main() {
	green := color.New(color.FgGreen)
	start := time.Now()

	green.Printf("Updating your repos %s at %v\n\n", repos, start)

	for _, repo := range repos {
		go runRepo(repo)
	}

	for counter < len(repos) {
		//blergh
	}

	duration := time.Since(start)
	green.Printf("\nDone updating. Time elapsed: %v\n", duration)
}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

// writeFile writes a file of given name to the output dir for project
func writeFile(project, projectPath *string, file string, content []byte) {
	path := fmt.Sprintf("%s/%s/%s", *projectPath, *project, file)
	err := os.WriteFile(path, content, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

// makeProject creates the stubbed files for a project
func makeProject(project, projectPath, builderType *string) {
	var b []byte
	// make dir from project name
	path := fmt.Sprintf("%s/%s", *projectPath, *project)
	err := os.MkdirAll(path, 0755)

	// continue if dir already exists
	if err != nil {
		log.Fatalln(err)
	}

	// create empty main.go
	writeFile(project, projectPath, "main.go", nil)

	// create go.mod
	version := strings.SplitN(strings.Replace(runtime.Version(), "go", "", 1), ".", 3)
	b = []byte(fmt.Sprintf("module %s\n\ngo %s", *project, version[0]+"."+version[1]))
	writeFile(project, projectPath, "go.mod", b)

	// create Dockerfile
	writeDocker(project, projectPath)

	// create Make/Task File
	writeBuildFIle(builderType, project, projectPath)

	// create .gitignore with project name added
	b = []byte(fmt.Sprintf("%s\n", *project))
	writeFile(project, projectPath, ".gitignore", b)
}

// flagCheck will print usage if no flags are passed in CLI
// or project name is missing
func flagCheck(project *string) {
	if flag.NFlag() == 0 || *project == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	// default output dir
	var defaultOutputDir = "output"

	// project flag
	builderType := flag.String("b", "task", "The `build file type of Make or Task to use `[make|task]. (OPTIONAL)")
	projectName := flag.String("n", "newproject", "The `name of the project` to template out. (REQUIRED)")
	projectPath := flag.String("p", "./output", "The `path to create the files in`. (OPTIONAL)")

	// parse inputted flags from CLI
	flag.Parse()

	// check minimum flags
	flagCheck(projectName)

	// if no project path specified create default output dir if missing
	if *projectPath == "" {
		// set projectPath to default
		*projectPath = defaultOutputDir
		err := os.Mkdir(defaultOutputDir, 0755)
		if err != nil && !os.IsExist(err) {
			log.Fatalln(err)
		}
	}

	makeProject(projectName, projectPath, builderType)
}

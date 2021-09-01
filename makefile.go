package main

import "strings"

const makeFileContent = `## Generic Makefile
.PHONY: clean
.DEFAULT_GOAL := build

build:
	go build -v ./...

clean:
	rm -f <project_name>
`

// writeDocker writes a Dockerfile to the project output dir from
// const dockerFileContent
func writeMake(project, projectPath *string) {
	var b []byte

	// replace <project_name> with project
	// create Dockerfile from const
	b = []byte(strings.Replace(makeFileContent, "<project_name>", *project, -1))
	writeFile(project, projectPath, "Makefile", b)
}
package main

import "strings"

const makeFileContent = `## Generic Makefile
.PHONY: clean
.DEFAULT_GOAL := build

build:
	docker-compose build

build_bin:
	go build -v ./...

push:
	docker-compose push

run:
	docker-compose up

clean:
	-rm <project_name>
	-docker-compose down
	-docker rmi <project_name>
`

const taskFileContent = `## Generic Taskfile
version: '3'

tasks:
  default:
    cmds:
      - task: run

  build:
    cmds:
      - docker-compose build

  build_bin:
    cmds:
      - go build -v ./...

  push:
    cmds:
      - docker-compose push

  run:
    cmds:
      - task: build
      - docker-compose up

  clean:
    cmds:
      - rm <project_name>
      - docker rmi -f <project_name>
    ignore_error: true
`

// writeBuildFIle writes a Make or Task file to the project output dir
func writeBuildFIle(builderType, project, projectPath *string) {
	var (
		b        []byte
		fileName string
	)

	// replace <project_name> with project
	// create Dockerfile from const
	switch *builderType {
	case "make":
		b = []byte(strings.Replace(makeFileContent, "<project_name>", *project, -1))
		fileName = "Makefile"
	default:
		b = []byte(strings.Replace(taskFileContent, "<project_name>", *project, -1))
		fileName = "Taskfile.yml"
	}
	writeFile(project, projectPath, fileName, b)
}

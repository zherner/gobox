package main

const dockerFileContent = `## Generic Go dockerfile
# Stage: build

# https://hub.docker.com/_/golang
FROM golang:alpine AS buildStage

ENV GOBIN=/go/bin
WORKDIR /go/src

COPY ./ ./

RUN go install -mod=vendor ./...

# Stage: final
FROM alpine:latest AS final

RUN apk add \
  curl

COPY --from=buildStage /go/bin/* /usr/local/bin/

WORKDIR /usr/local/bin

# Sets command to run the first executable in the current dir. This will be the just built golang binary
CMD find ./ -maxdepth 1 -perm -111 -type f -exec {} \;
`

const dockerComposeFileContent = `## Generic docker-compose
version: "3"

services:

  main:
    environment:
      ENV_VAR1: 'NOT_SET'
      ENV_VAR2: 'NOT_SET'
    build:
      context: .
    # command: |-
    #   requestbot -c 1 -u https://google.com
    image: IMAGE_URI/NAME
    ports:
      - "8080:8080"
`

// writeDocker writes a Dockerfile to the project output dir from
// const dockerFileContent
func writeDocker(project, projectPath *string) {
	var b []byte

	// create Dockerfile from const
	b = []byte(dockerFileContent)
	writeFile(project, projectPath, "Dockerfile", b)

    // create docker-compose from const
    b = []byte(dockerComposeFileContent)
    writeFile(project, projectPath, "docker-compose.yaml", b)
}

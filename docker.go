package main

const dockerFileContent = `## Generic Go dockerfile
# Stage: build
FROM golang:1.16-alpine AS buildStage

ENV GOBIN=/go/bin
WORKDIR /go/src

COPY ./ ./

RUN go install -mod=vendor ./...

# Stage: final
FROM alpine:3.14 AS final

RUN apk add \
  curl

COPY --from=buildStage /go/bin/* /usr/local/bin/

WORKDIR /usr/local/bin

#CMD [""]
`

// writeDocker writes a Dockerfile to the project output dir from
// const dockerFileContent
func writeDocker(project, projectPath *string) {
	var b []byte

	// create Dockerfile from const
	b = []byte(dockerFileContent)
	writeFile(project, projectPath, "Dockerfile", b)
}

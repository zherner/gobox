# gobox

A tool to template out common files in a Golang Docker project quickly.

# What is Gobox

A tool that creates a directory and some common needed files for a Golang Docker project like
- Dockerfile
- go.mod
- .gitignore
- Makefile

Gobox takes two input options
- `-n projectName` (Required)
- `-p /path/to/project/dir` (Optional. Default is current dir/output)
# How to use Gobox

First either download binary or build from source:

- Download the binary from the release page.

or

- Clone the repo: git clone https://github.com/zherner/gobox.git

- In the repo dir: `make`

- cd $GOPATH/bin

Then:

Create template project (the -p path/to/dir is optional): `gobox -n testProject -p /projects/`

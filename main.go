package main

import (
    "flag"
    "fmt"
    "log"
    "os"
)

var outputDir = "output"

// flagCheck will print usage if no flags are passed in CLI
func flagCheck(){
    if flag.NFlag() == 0 {
        flag.Usage()
        os.Exit(1)
    }
}

// makeProject
func makeProject(project *string) {
    // make dir from project name
    path := fmt.Sprintf("%s/%s", outputDir, *project)
    err := os.Mkdir(path, 0755)
    // continue if dir already exists
    if err != nil && !os.IsExist(err) {
        log.Fatalln(err)
    }

    // create empty main.go
    writeFile(*project, "main.go", nil)

    // create .gitignore with project name added
    b := []byte(*project)
    writeFile(*project, ".gitignore", b)
}

// writeFile
func writeFile(project, file string, content []byte) {
    path := fmt.Sprintf("%s/%s/%s", outputDir, project, file)
    err := os.WriteFile(path, content, 0755)
    if err != nil {
        log.Fatalln(err)
    }
}


func main() {
    // project flag
    project := flag.String("p", "", "The name of the proejct to template out.")

    // parse inputted flags from CLI
    flag.Parse()

    // check minimum flags
    flagCheck()

    //
    makeProject(project)
}

// Command driver runs the Go that codespec-write-go writes from codespec-read-ignite's document
// of ../../source/builds.ignite: it reads and ignores the harness trace and prints one line per
// case. The package ./builds holds export.go on disk; `go build -overlay overlay-read.json`
// supplies builds.go from the build directory the reader-builds task writes it to.
package main

import (
	"bufio"
	"io"
	"os"

	"codespec.example/buildsdriver/builds"
)

func main() {
	if _, err := io.ReadAll(os.Stdin); err != nil {
		os.Exit(2)
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, line := range builds.Lines() {
		out.WriteString(line + "\n")
	}
}

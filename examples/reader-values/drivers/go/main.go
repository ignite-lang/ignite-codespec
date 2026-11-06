// Command driver runs the Go that codespec-write-go writes from codespec-read-ignite's document
// of ../../source/values.ignite: it reads and ignores the harness trace and prints one line per
// case. The package ./values holds export.go on disk; `go build -overlay overlay-read.json`
// supplies values.go from the build directory the reader-values task writes it to.
package main

import (
	"bufio"
	"io"
	"os"

	"codespec.example/valuesdriver/values"
)

func main() {
	if _, err := io.ReadAll(os.Stdin); err != nil {
		os.Exit(2)
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, line := range values.Lines() {
		out.WriteString(line + "\n")
	}
}
